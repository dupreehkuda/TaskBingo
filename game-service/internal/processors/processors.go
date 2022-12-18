package processors

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/game-service/internal/interfaces"
)

type processor struct {
	userStorage i.UserDataClient
	taskStorage i.TaskDataClient
	logger      *zap.Logger
}

// New creates new instance of actions
func New(userStorage i.UserDataClient, taskStorage i.TaskDataClient, logger *zap.Logger) *processor {
	return &processor{
		userStorage: userStorage,
		taskStorage: taskStorage,
		logger:      logger,
	}
}
