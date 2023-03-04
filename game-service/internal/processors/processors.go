package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/taskClient"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/userClient"
)

// Processor is an interface for business-logic
type Processor interface {
	GetUserData(login string) (*models.UserAccountInfo, error)
	LoginUser(login, password string) (string, error)
	RegisterUser(creds *models.RegisterCredentials) (string, error)
	GetTaskPack(packID string) (*models.TaskPack, error)
	SetTaskPack(login string, pack *models.TaskPack) error
	LikeTaskPack(login, pack string) error
	DislikeTaskPack(login, pack string) error
	RateTaskPack(login, pack string) error
	UnrateTaskPack(login, pack string) error
	GetRatedPacks() (*[]models.TaskPack, error)
	GetAllUsers() (*[]models.User, error)
	AcceptFriend(login, person string) error
	DeleteFriend(login, person string) error
	RequestFriend(login, person string) error
	CreateGame(user1, user2, packID string) error
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
