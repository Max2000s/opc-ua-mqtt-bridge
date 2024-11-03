package handlers

import (
	"context"

	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/clients"
	"github.com/Max2000s/opc-ua-mqtt-bridge/pkg/config"
)

// The ready bit handler waits for the notification of a ready bit.
// Then it reads the defined nodes.
// When the nodes are read acknowledge bit is set.
type ReadyBitHandler struct {
}

func NewReadyBitHandler(handlerCfg config.HandlerConfig, clients *clients.Clients) (*ReadyBitHandler, error) {
	return nil, nil
}

func (readyBitHandler *ReadyBitHandler) Run(ctx context.Context) error {
	return nil
}

func (readyBitHandler *ReadyBitHandler) Name() string {
	return ""
}
