package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/clients"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/handlers"
)

type App struct {
	AppConfig  config.AppConfig
	AppClients *clients.Clients
	Handlers   []handlers.Handler
}

func NewApp(appConfig config.AppConfig) *App {
	log.Println("Initializing Application with config ...")

	clients, err := clients.InitializeClients(appConfig.Clients)
	if err != nil {
		log.Fatalf("Failed to initialize clients: %s", err)
	}

	handlers, err := InitializeHandlers(appConfig.Handlers, clients)
	if err != nil {
		log.Fatalf("Failed to initialize handlers: %s", err)
	}

	log.Printf("Initialized application!")

	return &App{
		AppConfig:  appConfig,
		AppClients: clients,
		Handlers:   handlers,
	}
}

func (app *App) Start() {
	log.Println("Starting application ...")

	if token := app.AppClients.MqttClient.Client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT: %v", token.Error())
	}
	log.Println("Connection to MQTT broker was successfull!")

	ctx := context.Background()
	if err := app.AppClients.OpcUaClient.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to OPC UA: %s", err)
	}

	defer func() {
		app.AppClients.OpcUaClient.Disconnect(ctx)
		app.AppClients.MqttClient.Disconnect(0) //quit directly
	}()

	app.StartHandlers()

	log.Println("Application stopped")

}

func InitializeHandlers(handlerConfigs []config.HandlerConfig, clients *clients.Clients) ([]handlers.Handler, error) {
	var handlerList []handlers.Handler
	for _, handlerCfg := range handlerConfigs {
		var handler handlers.Handler
		var err error

		switch handlerCfg.Type {
		case "ReadyBitHandler":
			handler, err = handlers.NewReadyBitHandler(handlerCfg, clients)
		default:
			log.Printf("Unknown handler type '%s' will be skipped", handlerCfg.Type)
			continue
		}

		if err != nil {
			log.Printf("Error while creaating handler type '%s': %s", handlerCfg.Type, err)
			continue
		}
		handlerList = append(handlerList, handler)
	}
	return nil, nil
}

func (app *App) StartHandlers() {

	// create waitgroup and context
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start goroutines
	for _, handler := range app.Handlers {
		wg.Add(1)
		go func(h handlers.Handler) {
			defer wg.Done()
			if err := h.Run(ctx); err != nil {
				log.Printf("Handler %s exited with error: %v", h.Name(), err)
			}
		}(handler)
	}

	// Handle shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()

	wg.Wait()
}
