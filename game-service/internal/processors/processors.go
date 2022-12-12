package processors

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/game-service/internal/interfaces"
)

type processor struct {
	userStorage i.UserDataClient
	logger      *zap.Logger
}

// New creates new instance of actions
func New(userStorage i.UserDataClient, logger *zap.Logger) *processor {
	return &processor{
		userStorage: userStorage,
		logger:      logger,
	}
}
