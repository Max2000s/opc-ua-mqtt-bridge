package opcuaclient

import (
	"context"
	"fmt"

	"github.com/gopcua/opcua/ua"
)

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
