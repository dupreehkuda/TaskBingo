package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/taskClient"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/userClient"
)

// Processor is an interface for business-logic
type Processor interface {
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
}

// processor provides service's business-logic
type processor struct {
	userStorage userClient.UserDataClient
	taskStorage taskClient.TaskDataClient
	logger      *zap.Logger
}

// New creates new instance of processor
func New(userStorage userClient.UserDataClient, taskStorage taskClient.TaskDataClient, logger *zap.Logger) *processor {
	return &processor{
		userStorage: userStorage,
		taskStorage: taskStorage,
		logger:      logger,
	}
}
