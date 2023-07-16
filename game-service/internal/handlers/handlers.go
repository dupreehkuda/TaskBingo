package handlers

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// Service is an interface for business-logic
type Service interface {
	GetUserData(ctx context.Context, userID string) (*models.UserAccountInfo, error)
	LoginUser(ctx context.Context, username, password string) (string, error)
	RegisterUser(ctx context.Context, credits *models.RegisterCredentials) (string, error)

	GetTaskPack(ctx context.Context, packID string) (*models.TaskPack, error)
	SetTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error
	LikeTaskPack(ctx context.Context, userID, pack string) error
	DislikeTaskPack(ctx context.Context, userID, pack string) error
	RateTaskPack(ctx context.Context, userID, pack string) error
	UnrateTaskPack(ctx context.Context, userID, pack string) error
	GetRatedPacks(ctx context.Context) (*models.Packs, error)

	GetAllUsers(ctx context.Context) (*models.Users, error)
	AcceptFriend(ctx context.Context, userID, friendID string) error
	DeleteFriend(ctx context.Context, userID, friendID string) error
	RequestFriend(ctx context.Context, userID, friendID string) error

	CreateGame(ctx context.Context, userID, opponentID, packID string) (*models.GameShort, error)
	GetGame(ctx context.Context, gameID string) (*models.Game, error)
	AcceptGame(ctx context.Context, userID, gameID string) error
	DeleteGame(ctx context.Context, userID, gameID string) error
	GetRoom(ctx context.Context, gameID string) (*models.Room, error)
	UpdateGame(ctx context.Context, room *models.Room, action *models.GameAction) (*models.GameUpdate, error)
}

// Handlers provides access to service
type handlers struct {
	service Service
	hub     *models.GameHub
	domain  string
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(service Service, domain string, logger *zap.Logger) *handlers {
	hub := &models.GameHub{
		Rooms: make(map[string]*models.Room),
	}

	return &handlers{
		service: service,
		hub:     hub,
		domain:  domain,
		logger:  logger,
	}
}

// UUIDCheck checks all request ids to be sure they're not empty and correct uuids
func UUIDCheck(uuids ...string) error {
	for _, id := range uuids {
		if id == "" {
			return errs.ErrEmptyRequest
		}

		_, err := uuid.Parse(id)
		if err != nil {
			return err
		}
	}

	return nil
}
