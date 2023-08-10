package handlers

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// Handler is interface for handlers
type Handler interface {
	api.BingoDataServer
}

// Service is interface for business-logic
type Service interface {
	RegisterUser(ctx context.Context, username, password, email, city string) (string, error)
	LoginUser(ctx context.Context, username, password string) (string, error)
	GetUserData(ctx context.Context, userID string) (*models.GetUserDataResponse, error)

	GetRatedPacks(ctx context.Context) (*[]models.TaskPack, error)
	AddTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error
	GetTaskPacks(ctx context.Context, packIDs ...string) (*[]models.TaskPack, error)

	LikePack(ctx context.Context, userID, pack string, inc int) error
	RatePack(ctx context.Context, userID, pack string, inc int) error

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

// Handlers provides access to service
type Handlers struct {
	service Service
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(service Service, logger *zap.Logger) *Handlers {
	return &Handlers{
		service: service,
		logger:  logger,
	}
}
