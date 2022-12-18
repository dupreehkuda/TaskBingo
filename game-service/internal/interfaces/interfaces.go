package interfaces

import (
	"net/http"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

type Handlers interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)
	GetTaskPack(w http.ResponseWriter, r *http.Request)
	SetTaskPack(w http.ResponseWriter, r *http.Request)
}

type UserDataClient interface {
	GetUserData(login string) (*models.Response, error)
	RegisterUser(login, email, password string) error
	LoginUser(login, password string) error
}

type TaskDataClient interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
}

// Middleware implement an interface for middleware layer
type Middleware interface {
	CheckToken(next http.Handler) http.Handler
}

type Processor interface {
	GetUserData(login string) (*models.Response, error)
	LoginUser(login, password string) (string, error)
	RegisterUser(login, email, password string) (string, error)
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
}
