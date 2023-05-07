package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// Repository is interface for repository
type Repository interface {
	GetUserData(ctx context.Context, userID string) (*models.GetUserDataResponse, error)
	CheckDuplicateUser(ctx context.Context, username, email string) (bool, error)
	CreateUser(ctx context.Context, userID, username, email, passwordHash, passwordSalt, city string) error
	LoginUser(ctx context.Context, username string) (*models.LoginUserResponse, error)

	GetRatedPacks(ctx context.Context) ([]string, error)
	LikePack(ctx context.Context, userID, pack string, inc int) error
	RatePack(ctx context.Context, userID, pack string, inc int) error
	AssignNewPack(ctx context.Context, userID, packID string, packName string) error

	GetAllUsers(ctx context.Context) (*[]models.AllUsers, error)
	AcceptFriend(ctx context.Context, userID, friendID string) error
	DeleteFriend(ctx context.Context, userID, friendID string) error
	RequestFriend(ctx context.Context, userID, friendID string) error

	CreateGame(ctx context.Context, game *models.Game) error
	AcceptGame(ctx context.Context, userID, gameID string) error
	DeleteGame(ctx context.Context, userID, gameID string) error
	GetGame(ctx context.Context, gameID string) (*models.Game, error)
	AchieveGame(ctx context.Context, game *models.Game) error
}

// service provides service's business logic
type service struct {
	repository Repository
	logger     *zap.Logger
}

// New creates new instance of actions
func New(repo Repository, logger *zap.Logger) *service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}
