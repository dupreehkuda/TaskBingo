package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// Repository is interface for repository
type Repository interface {
	GetTaskPack(taskId string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
	CheckPackExistence(taskId string) (bool, error)
}

// service provides service's business logic
type service struct {
	repository Repository
	logger     *zap.Logger
}

// New creates new instance of service
func New(repository Repository, logger *zap.Logger) *service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}
