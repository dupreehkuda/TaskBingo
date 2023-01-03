package interfaces

import (
	"net/http"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// Handlers is an interface for handlers
type Handlers interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)
	GetTaskPack(w http.ResponseWriter, r *http.Request)
	SetTaskPack(w http.ResponseWriter, r *http.Request)
}

// UserDataClient is an interface for user data service
type UserDataClient interface {
	GetUserData(login string) (*models.UserAccountInfoResponse, error)
	RegisterUser(login, email, password string) error
	LoginUser(login, password string) error
}

// TaskDataClient is an interface for task data service
type TaskDataClient interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
	GetFavouritePacks(packIDs []string) (*[]models.TaskPack, error)
}

// Middleware is an interface for middleware layer
type Middleware interface {
	CheckToken(next http.Handler) http.Handler
}

// Processor is an interface for business-logic
type Processor interface {
	GetUserData(login string) (*models.UserAccountInfo, error)
	LoginUser(login, password string) (string, error)
	RegisterUser(login, email, password string) (string, error)
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
}
