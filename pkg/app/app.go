package app

import (
	"log"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/opcuaclient"
)

type App struct {
	AppConfig   config.AppConfig
	OpcUaClient *opcuaclient.OpcUaClient
	MqttClient  string
}

func NewApp(appConfig config.AppConfig) *App {
	log.Println("initializing Application with config ...")

	opcuaClient, err := opcuaclient.NewClient(&appConfig.OpcUA)
	if err != nil {
		log.Fatalf("failed to create OPC UA client, will exit now with error: %s", err)
	}

	return &App{
		AppConfig:   appConfig,
		OpcUaClient: opcuaClient,
		MqttClient:  "nil",
	}
}
