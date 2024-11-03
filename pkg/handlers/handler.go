package handlers

import (
	"context"
)

// general interface description of a handler
type Handler interface {
	Run(ctx context.Context) error
	Name() string
}
