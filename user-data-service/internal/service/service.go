package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/repository"
)

// Service is interface for business-logic
type Service interface {
	RegisterUser(username, password, email, city string) (string, error)
	LoginUser(username, password string) (string, error)
	GetUserData(userID string) (*models.GetUserDataResponse, error)

	GetRatedPacks() ([]string, error)
	LikePack(userID, pack string, inc int) error
	RatePack(userID, pack string, inc int) error
	AssignNewPack(userID, packID string, packName string) error

	GetAllUsers() (*[]models.AllUsers, error)
	AcceptFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
	RequestFriend(userID, friendID string) error

	CreateGame(game *models.Game) error
	AcceptGame(userID, gameID string) error
	DeleteGame(userID, gameID string) error
}

// service provides service's business logic
type service struct {
	repository repository.Repository
	logger     *zap.Logger
}

// New creates new instance of actions
func New(repo repository.Repository, logger *zap.Logger) *service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}
