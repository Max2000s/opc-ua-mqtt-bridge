package main

import (
	"log"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/app"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
)

func main() {

	CONFIG_FILE := "config/config.ylaml"

	log.Printf("Start reading config file %s", CONFIG_FILE)
	appConfig, err := config.LoadConfig(CONFIG_FILE)
	if err != nil {
		log.Fatalf("There was a problem reading the config file: %s", err)
	}

	app := app.NewApp(*appConfig)
	app.Start()

}
