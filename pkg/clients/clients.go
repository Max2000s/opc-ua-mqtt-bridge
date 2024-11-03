package clients

import (
	"fmt"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
)

// posible to add more clients
type Clients struct {
	OpcUaClient *OpcUaClient
	MqttClient  *MqttClient
}

func InitializeClients(clientsConfig config.ClientsConfig) (*Clients, error) {

	opcuaClient, err := NewOpcUaClient(&clientsConfig.OpcUA)
	if err != nil {
		return nil, fmt.Errorf("failed to create OPC UA client: %s", err)
	}

	mqttClient, err := NewMqttClient(clientsConfig.MQTT)
	if err != nil {
		return nil, fmt.Errorf("failed to create MQTT client: %s", err)
	}

	return &Clients{
		OpcUaClient: opcuaClient,
		MqttClient:  mqttClient,
	}, nil
}
