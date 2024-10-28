package app

import (
	"context"
	"log"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/mqttclient"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/opcuaclient"
)

type App struct {
	AppConfig   config.AppConfig
	OpcUaClient *opcuaclient.OpcUaClient
	MqttClient  *mqttclient.MQTTClient
}

func NewApp(appConfig config.AppConfig) *App {
	log.Println("Initializing Application with config ...")

	opcuaClient, err := opcuaclient.NewOpcUaClient(&appConfig.OpcUA)
	if err != nil {
		log.Fatalf("failed to create OPC UA client, will exit now with error: %s", err)
	}

	mqttClient, err := mqttclient.NewMqttClient(appConfig.MQTT)
	if err != nil {
		log.Fatalf("failed to create MQTT client, will exit now with error: %s", err)
	}

	return &App{
		AppConfig:   appConfig,
		OpcUaClient: opcuaClient,
		MqttClient:  mqttClient,
	}
}

func (app *App) Start() {
	log.Println("Starting application ...")

	// connect to OPCUA
	ctx := context.Background()
	if err := app.OpcUaClient.Connect(ctx); err != nil {
		log.Println("Failed to connect: %v\n", err)
		return
	}

	defer app.OpcUaClient.Disconnect(ctx)
}
