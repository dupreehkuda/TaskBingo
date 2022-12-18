package processors

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/task-data-service/internal/interfaces"
)

// Processor provides service's business logic
type processor struct {
	storage i.Stored
	logger  *zap.Logger
}

// New creates new instance of actions
func New(userStorage i.Stored, logger *zap.Logger) *processor {
	return &processor{
		storage: userStorage,
		logger:  logger,
	}
}
