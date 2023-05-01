package handlers

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/service"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// Handler is interface for handlers
type Handler interface {
	api.UsersServer
}

// Handlers provides access to service
type Handlers struct {
	service service.Service
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(service service.Service, logger *zap.Logger) *Handlers {
	return &Handlers{
		service: service,
		logger:  logger,
	}
}
