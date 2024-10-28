package opcuaclient

import (
	"context"
	"fmt"
	"time"

	"github.com/gopcua/opcua"
)

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
