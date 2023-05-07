package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// UserRepository is an interface for user data service
type UserRepository interface {
	GetUserData(userID string) (*models.UserAccountInfoResponse, error)
	RegisterUser(credits *models.RegisterCredentials) (string, error)
	LoginUser(username, password string) (string, error)

	GetRatedPacks() ([]string, error)
	LikeTaskPack(userID, pack string) error
	DislikeTaskPack(userID, pack string) error
	RateTaskPack(userID, pack string) error
	UnrateTaskPack(userID, pack string) error
	AssignNewPack(userID, packID, packName string) error

	GetAllUsers() (*[]models.User, error)
	AcceptFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
	RequestFriend(userID, friendID string) error

	CreateGame(game *models.Game) error
	AcceptGame(userID, packID string) error
	DeleteGame(userID, packID string) error
	GetGame(gameID string) (*models.Game, error)
	AchieveGame(game *models.Game) error
}

// TaskRepository is an interface for task data service
type TaskRepository interface {
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(pack *models.TaskPack) error
	GetMultiplePacks(packIDs []string) (*[]models.TaskPack, error)
}

// service provides service's business-logic
type service struct {
	userRepository UserRepository
	taskRepository TaskRepository
	logger         *zap.Logger
}

// New creates new instance of processor
func New(userRepository UserRepository, taskRepository TaskRepository, logger *zap.Logger) *service {
	return &service{
		userRepository: userRepository,
		taskRepository: taskRepository,
		logger:         logger,
	}
}
