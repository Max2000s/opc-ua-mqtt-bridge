package config

type AppConfig struct {
	OpcUA OpcUaClientConfig `json:"opcua_client"`
	MQTT  MqttConfig        `json:"mqtt"`
}
