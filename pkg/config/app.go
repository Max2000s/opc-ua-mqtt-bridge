package config

type AppConfig struct {
	Clients  ClientsConfig   `yaml:"clients"`
	Handlers []HandlerConfig `yaml:"handlers"`
}

type ClientsConfig struct {
	OpcUA OpcUaClientConfig `yaml:"opcua_client"`
	MQTT  MqttConfig        `yaml:"mqtt"`
}

type HandlerConfig struct {
	Type       string                 `yaml:"type"`
	Name       string                 `yaml:"name"`
	Parameters map[string]interface{} `yaml:"parameters"`
}
