package config

type AppConfig struct {
	OpcUA    OpcUaClientConfig `yaml:"opcua_client"`
	MQTT     MqttConfig        `yaml:"mqtt"`
	Handlers []HandlerConfig   `yaml:"handlers"`
}

type HandlerConfig struct {
	Type       string                 `yaml:"type"`
	Name       string                 `yaml:"name"`
	Parameters map[string]interface{} `yaml:"parameters"`
}
