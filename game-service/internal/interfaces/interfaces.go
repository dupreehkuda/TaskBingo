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
	LikeTaskPack(w http.ResponseWriter, r *http.Request)
	DislikeTaskPack(w http.ResponseWriter, r *http.Request)
	RateTaskPack(w http.ResponseWriter, r *http.Request)
	UnrateTaskPack(w http.ResponseWriter, r *http.Request)
	GetRatedPacks(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	RequestFriend(w http.ResponseWriter, r *http.Request)
	AcceptFriend(w http.ResponseWriter, r *http.Request)
	DeleteFriend(w http.ResponseWriter, r *http.Request)
	CreateGame(w http.ResponseWriter, r *http.Request)
}

// UserDataClient is an interface for user data service
type UserDataClient interface {
	GetUserData(login string) (*models.UserAccountInfoResponse, error)
	RegisterUser(creds *models.RegisterCredentials) error
	LoginUser(login, password string) error
	GetRatedPacks() ([]string, error)
	LikeTaskPack(login, pack string) error
	DislikeTaskPack(login, pack string) error
	RateTaskPack(login, pack string) error
	UnrateTaskPack(login, pack string) error
	AssignNewPack(login, pack string) error
	GetAllUsers() (*[]models.User, error)
	AcceptFriend(login, person string) error
	DeleteFriend(login, person string) error
	RequestFriend(login, person string) error
	CreateGame(game *models.Game) error
}

// TaskDataClient is an interface for task data service
type TaskDataClient interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(packIDs []string) (*[]models.TaskPack, error)
}

// Middleware is an interface for middleware layer
type Middleware interface {
	CheckToken(next http.Handler) http.Handler
}

// Processor is an interface for business-logic
type Processor interface {
	GetUserData(login string) (*models.UserAccountInfo, error)
	LoginUser(login, password string) (string, error)
	RegisterUser(creds *models.RegisterCredentials) (string, error)
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(login string, pack *models.TaskPack) error
	LikeTaskPack(login, pack string) error
	DislikeTaskPack(login, pack string) error
	RateTaskPack(login, pack string) error
	UnrateTaskPack(login, pack string) error
	GetRatedPacks() (*[]models.TaskPack, error)
	GetAllUsers() (*[]models.User, error)
	AcceptFriend(login, person string) error
	DeleteFriend(login, person string) error
	RequestFriend(login, person string) error
	CreateGame(user1, user2, packID string) error
}
