package config

type MqttConfig struct {
	Broker   string `json:"broker"`
	ClientID string `json:"client_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Topic    string `json:"topic"`
}
