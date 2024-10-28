package opcuaclient

import (
	"context"
	"fmt"

	"github.com/gopcua/opcua/ua"
)

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
