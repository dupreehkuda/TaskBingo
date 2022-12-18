package handlers

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/task-data-service/internal/interfaces"
)

// Handlers provides access to service
type Handlers struct {
	processor i.Processor
	logger    *zap.Logger
}

// New creates new instance of handlers
func New(storage i.Processor, logger *zap.Logger) *Handlers {
	return &Handlers{
		processor: storage,
		logger:    logger,
	}
}
