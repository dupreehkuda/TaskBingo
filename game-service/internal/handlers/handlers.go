package handlers

import (
	"github.com/google/uuid"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// Service is an interface for business-logic
type Service interface {
	GetUserData(userID string) (*models.UserAccountInfo, error)
	LoginUser(username, password string) (string, error)
	RegisterUser(credits *models.RegisterCredentials) (string, error)

	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(userID string, pack *models.TaskPack) error
	LikeTaskPack(userID, pack string) error
	DislikeTaskPack(userID, pack string) error
	RateTaskPack(userID, pack string) error
	UnrateTaskPack(userID, pack string) error
	GetRatedPacks() (*[]models.TaskPack, error)

	GetAllUsers() (*[]models.User, error)
	AcceptFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
	RequestFriend(userID, friendID string) error

	CreateGame(userID, opponentID, packID string) error
	AcceptGame(userID, gameID string) error
	DeleteGame(userID, gameID string) error
	GetRoom(gameID string) (*models.Room, error)
	UpdateGame(room *models.Room, action *models.GameAction) (*models.GameUpdate, error)
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
