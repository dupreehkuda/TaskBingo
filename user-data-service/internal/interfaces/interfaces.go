package interfaces

import (
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// Handlers is interface for handlers
type Handlers interface {
	api.UsersServer
}

// Processor is interface for business-logic
type Processor interface {
	RegisterUser(login, password, email, city string) error
	LoginUser(login, password string) error
	GetUserData(userId string) (*models.GetUserDataResponse, error)
}

// Stored is interface for storage
type Stored interface {
	GetUserData(login string) (models.GetUserDataResponse, error)
	CheckDuplicateUser(login string) (bool, error)
	CreateUser(login, email, passwordHash, passwordSalt, city string) error
	LoginUser(login string) (*models.LoginUserResponse, error)
}
