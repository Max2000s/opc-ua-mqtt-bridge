package opcuaclient

import (
	"context"
	"fmt"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
	"github.com/gopcua/opcua"
)

type OpcUaClient struct {
	config *config.OpcUaClientConfig
	client *opcua.Client
}

// NewClient creates a new OPC UA client with the given configuration
func NewOpcUaClient(config *config.OpcUaClientConfig) (*OpcUaClient, error) {
	// Create OPC UA client options
	opts := []opcua.Option{
		opcua.SecurityPolicy(config.SecurityPolicy),
		opcua.SecurityModeString(config.SecurityMode),
	}

	// Add authentication methods if provided
	if config.Username != "" && config.Password != "" {
		opts = append(opts, opcua.AuthUsername(config.Username, config.Password))
	}

	// Add certificates if provided
	if config.Certificate != "" && config.PrivateKey != "" {
		//opts = append(opts, opcua.AuthCertificate())
		//cert, err := opcua.NewClientCertificate(config.Certificate, config.PrivateKey)
		//if err != nil {
		//    return nil, fmt.Errorf("Failed to load certificates: %w", err)
		//}
		//opts = append(opts, cert)
		return nil, fmt.Errorf("failed to load certificates: not implemented yet")
	}

	// Create the OPC UA client
	opcuaClient, err := opcua.NewClient(config.Endpoint, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpcUa client: %s", err)
	}

	return &OpcUaClient{
		config: config,
		client: opcuaClient,
	}, nil
}

func (c *OpcUaClient) Connect(ctx context.Context) error {
	if err := c.client.Connect(ctx); err != nil {
		return fmt.Errorf("failed to connect to OPC UA server: %w", err)
	}
	return nil
}

func (c *OpcUaClient) Disconnect(ctx context.Context) error {
	return c.client.Close(ctx)
}
