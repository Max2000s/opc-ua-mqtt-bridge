package clients

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
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
	endpoints, err := opcua.GetEndpoints(context.Background(), c.config.Endpoint)
	if err != nil {
		log.Fatalf("Failed to get endpoints: %v", err)
	}

	for _, ep := range endpoints {
		log.Printf("Endpoint URL: %s", ep.EndpointURL)
		// You can look for the endpoint that matches your criteria
		// For example, matching the URL that ends with "/milosb"
	}

	if err := c.client.Connect(ctx); err != nil {
		return fmt.Errorf("failed to connect to OPC UA server: %w", err)
	}
	return nil
}

func (c *OpcUaClient) Disconnect(ctx context.Context) error {
	return c.client.Close(ctx)
}

// ReadValue reads a value from the given node ID
func (c *OpcUaClient) ReadValue(nodeID *ua.NodeID) (*ua.DataValue, error) {
	ctx := context.Background()
	req := &ua.ReadRequest{
		NodesToRead: []*ua.ReadValueID{
			{
				NodeID:      nodeID,
				AttributeID: ua.AttributeIDValue,
			},
		},
	}
	resp, err := c.client.Read(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Results[0].Status != ua.StatusOK {
		return nil, fmt.Errorf("read failed with status: %v", resp.Results[0].Status)
	}
	return resp.Results[0], nil
}

// WriteValue writes a value to the given node ID
func (c *OpcUaClient) WriteValue(nodeID *ua.NodeID, value interface{}) error {
	v, err := ua.NewVariant(value)
	if err != nil {
		return err
	}
	ctx := context.Background()
	req := &ua.WriteRequest{
		NodesToWrite: []*ua.WriteValue{
			{
				NodeID:      nodeID,
				AttributeID: ua.AttributeIDValue,
				Value: &ua.DataValue{
					Value: v,
				},
			},
		},
	}
	resp, err := c.client.Write(ctx, req)
	if err != nil {
		return err
	}
	if resp.Results[0] != ua.StatusOK {
		return fmt.Errorf("write failed with status: %v", resp.Results[0])
	}
	return nil
}

// Subscription represents a subscription to OPC UA monitored items
type Subscription struct {
	client       *OpcUaClient
	subscription *opcua.Subscription
}

// CreateSubscription creates a new subscription with the given parameters
func (c *OpcUaClient) CreateSubscription(interval time.Duration) (*Subscription, error) {
	ctx := context.Background()
	sub, err := c.client.Subscribe(ctx, &opcua.SubscriptionParameters{
		Interval: interval,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}
	return &Subscription{
		client:       c,
		subscription: sub,
	}, nil
}

/*
// Subscribe adds monitored items to the subscription
func (s *Subscription) Subscribe(nodeIDs []*ua.NodeID, handler opcua.PublishNotificationHandler) error {
	var requests []*ua.MonitoredItemCreateRequest
	for i, nodeID := range nodeIDs {
		handle := uint32(i + 1)
		request := opcua.NewMonitoredItemCreateRequestWithDefaults(nodeID, ua.AttributeIDValue, handle)
		requests = append(requests, request)
	}
	err := s.subscription.SetPublishingMode(true)
	if err != nil {
		return fmt.Errorf("failed to set publishing mode: %w", err)
	}
	_, err = s.subscription.Monitor(ua.TimestampsToReturnBoth, requests...)
	if err != nil {
		return fmt.Errorf("failed to monitor items: %w", err)
	}
	s.subscription.AddNotificationHandler(handler)
	return nil
}*/

// Cancel cancels the subscription
func (s *Subscription) Cancel(ctx context.Context) error {
	return s.subscription.Cancel(ctx)
}
