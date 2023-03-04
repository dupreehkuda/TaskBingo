package handlers

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/processors"
	api "github.com/dupreehkuda/TaskBingo/task-data-service/pkg/api"
)

// Handler is interface for handlers
type Handler interface {
	api.TasksServer
}

// Handlers provides access to service
type Handlers struct {
	processor processors.Processor
	logger    *zap.Logger
}

// New creates new instance of handlers
func New(storage processors.Processor, logger *zap.Logger) *Handlers {
	return &Handlers{
		processor: storage,
		logger:    logger,
	}
}
