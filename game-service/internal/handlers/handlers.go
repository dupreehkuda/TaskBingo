package handlers

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/game-service/internal/interfaces"
)

// Handlers provides access to service
type handlers struct {
	processor i.Processor
	logger    *zap.Logger
}

// New creates new instance of handlers
func New(processor i.Processor, logger *zap.Logger) *handlers {
	return &handlers{
		processor: processor,
		logger:    logger,
	}
}
