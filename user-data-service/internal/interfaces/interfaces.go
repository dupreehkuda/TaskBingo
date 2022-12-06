package interfaces

import (
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

type Handlers interface {
	api.UsersServer
}

type Stored interface {
	Ping(userID string) (models.Response, error)
}
