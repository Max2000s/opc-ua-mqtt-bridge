package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/handlers"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/mqttclient"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/opcuaclient"
)

type App struct {
	AppConfig   config.AppConfig
	OpcUaClient *opcuaclient.OpcUaClient
	MqttClient  *mqttclient.MQTTClient
	Handlers    []handlers.Handler
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
		log.Fatalf("Failed to connect: %s", err)
	}

	// connect to MQTT

	defer func() {
		app.OpcUaClient.Disconnect(ctx)
		app.MqttClient.Disconnect(ctx)
	}()

	var wg sync.WaitGroup

	if err := app.InitializeHandlers(); err != nil {
		log.Fatalf("Failed to initialize handlers: %s", err)
	}

	app.StartHandlers()

	// Wait for an interrupt signal to gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Wait for all handlers to finish
	wg.Wait()
	log.Println("Application stopped")
}

func (app *App) InitializeHandlers() error {
	for _, handlerCfg := range app.AppConfig.Handlers {
		var handler handlers.Handler
		var err error

		switch handlerCfg.Type {
		case "ReadyBitHandler":
			handler, err = handlers.NewReadyBitHandler()
		default:
			log.Printf("Unknown handler type: %s", handlerCfg.Type)
			continue
		}

		if err != nil {
			log.Printf("Error while creaating handler type '%s': %s", handlerCfg.Type, err)
			continue
		}

		app.Handlers = append(app.Handlers, handler)
	}
	return nil
}

func (app *App) StartHandlers() error {
	// fill handler start here

	return nil
}
