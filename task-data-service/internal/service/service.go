package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/repository"
)

// Service is interface for business-logic
type Service interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	AddTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(ids []string) (*[]models.TaskPack, error)
}

// service provides service's business logic
type service struct {
	repository repository.Repository
	logger     *zap.Logger
}

// New creates new instance of service
func New(repository repository.Repository, logger *zap.Logger) *service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}
