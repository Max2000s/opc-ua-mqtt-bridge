package mqttclient

import (
	"log"
	"os"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	Client mqtt.Client
	Config config.MqttConfig
}

func NewMQTTClient(cfg config.MqttConfig) *MQTTClient {
	opts := mqtt.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientID).
		SetUsername(cfg.Username).
		SetPassword(cfg.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
		os.Exit(1)
	}

	log.Printf("Connected to MQTT broker '%s'.", cfg.Broker)

	return &MQTTClient{
		Client: client,
		Config: cfg,
	}
}

func (mc *MQTTClient) Disconnect(quiesce uint) {
	mc.Client.Disconnect(quiesce)
}

func (c *MQTTClient) Subscribe(topic string, callback mqtt.MessageHandler) {
	if token := c.Client.Subscribe(topic, 1, callback); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, token.Error())
	}
}

func ListenOnTopic(topic string, mqttClient *MQTTClient, rawMessages chan<- mqtt.Message) {
	log.Printf("MQTT Client will listen for messages on the topic '%s'.", topic)
	mqttClient.Subscribe(topic, func(client mqtt.Client, msg mqtt.Message) {
		rawMessages <- msg
	})
	select {} // keep function alive
}
