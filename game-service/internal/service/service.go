package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// Repository is an interface for user data service
type Repository interface {
	GetUserData(ctx context.Context, userID string) (*models.UserAccountInfo, error)
	RegisterUser(ctx context.Context, credits *models.RegisterCredentials) (string, error)
	LoginUser(ctx context.Context, username, password string) (string, error)

	GetRatedPacks(ctx context.Context) (*models.Packs, error)
	LikeTaskPack(ctx context.Context, userID, pack string) error
	DislikeTaskPack(ctx context.Context, userID, pack string) error
	RateTaskPack(ctx context.Context, userID, pack string) error
	UnrateTaskPack(ctx context.Context, userID, pack string) error
	SetNewTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error
	GetTaskPacks(ctx context.Context, packIDs ...string) (*models.Packs, error)

	GetAllUsers(ctx context.Context) (*models.Users, error)
	AcceptFriend(ctx context.Context, userID, friendID string) error
	DeleteFriend(ctx context.Context, userID, friendID string) error
	RequestFriend(ctx context.Context, userID, friendID string) error

	CreateGame(ctx context.Context, game *models.Game) error
	AcceptGame(ctx context.Context, userID, packID string) error
	DeleteGame(ctx context.Context, userID, packID string) error
	GetGame(ctx context.Context, gameID string) (*models.Game, error)
	AchieveGame(ctx context.Context, game *models.Game) error
}

// TaskRepository is no more in use as we stopped using redis for storing tasks
//
//// TaskRepository is an interface for task data service
//type TaskRepository interface {
//	GetTaskPack(ctx context.Context, packID string) (*models.TaskPack, error)
//	SetTaskPack(ctx context.Context, pack *models.TaskPack) error
//	GetMultiplePacks(ctx context.Context, packIDs []string) (*models.Packs, error)
//}

// service provides service's business-logic
type service struct {
	repository Repository
	logger     *zap.Logger
}

// New creates new instance of processor
func New(Repository Repository, logger *zap.Logger) *service {
	return &service{
		repository: Repository,
		logger:     logger,
	}
}
