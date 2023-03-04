package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/storage"
)

// Processor is interface for business-logic
type Processor interface {
	GetTaskPack(taskId string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
}

// Processor provides service's business logic
type processor struct {
	storage storage.Stored
	logger  *zap.Logger
}

// New creates new instance of actions
func New(userStorage storage.Stored, logger *zap.Logger) *processor {
	return &processor{
		storage: userStorage,
		logger:  logger,
	}
}
