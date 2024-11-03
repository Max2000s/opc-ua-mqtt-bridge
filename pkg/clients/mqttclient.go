package clients

import (
	"log"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	Client mqtt.Client
	Config config.MqttConfig
}

func NewMqttClient(cfg config.MqttConfig) (*MqttClient, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientID).
		SetUsername(cfg.Username).
		SetPassword(cfg.Password)

	client := mqtt.NewClient(opts)

	return &MqttClient{
		Client: client,
		Config: cfg,
	}, nil
}

func (mc *MqttClient) Disconnect(quiesce uint) {
	mc.Client.Disconnect(quiesce)
}

func (c *MqttClient) Subscribe(topic string, callback mqtt.MessageHandler) {
	if token := c.Client.Subscribe(topic, 1, callback); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, token.Error())
	}
}

func ListenOnTopic(topic string, mqttClient *MqttClient, rawMessages chan<- mqtt.Message) {
	log.Printf("MQTT Client will listen for messages on the topic '%s'.", topic)
	mqttClient.Subscribe(topic, func(client mqtt.Client, msg mqtt.Message) {
		rawMessages <- msg
	})
	select {} // keep function alive
}
