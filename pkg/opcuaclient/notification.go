package opcuaclient

import (
	"fmt"

	"github.com/gopcua/opcua"
)

// ExampleNotificationHandler handles incoming notifications
func ExampleNotificationHandler(data *opcua.PublishNotificationData) {
	fmt.Print(data.Value)
}
