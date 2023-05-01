package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/taskRepository"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/userRepository"
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
}

// service provides service's business-logic
type service struct {
	userRepository userRepository.UserRepository
	taskRepository taskRepository.TaskRepository
	logger         *zap.Logger
}

// New creates new instance of processor
func New(userRepository userRepository.UserRepository, taskRepository taskRepository.TaskRepository, logger *zap.Logger) *service {
	return &service{
		userRepository: userRepository,
		taskRepository: taskRepository,
		logger:         logger,
	}
}
