package interfaces

import (
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/task-data-service/pkg/api"
)

// Handlers is interface for handlers
type Handlers interface {
	api.TasksServer
}

// Processor is interface for business-logic
type Processor interface {
	GetTaskPack(taskId string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
}

// Stored is interface for storage
type Stored interface {
	GetTaskPack(taskId string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
}
