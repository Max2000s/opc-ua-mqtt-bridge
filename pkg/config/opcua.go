package config

type OpcUaClientConfig struct {
	Endpoint       string
	SecurityPolicy string
	SecurityMode   string
	Certificate    string
	PrivateKey     string
	Username       string
	Password       string
}
