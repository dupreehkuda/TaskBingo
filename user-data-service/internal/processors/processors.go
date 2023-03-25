package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/storage"
)

// Processor is interface for business-logic
type Processor interface {
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
}

// Processor provides service's business logic
type processor struct {
	storage storage.Stored
	logger  *zap.Logger
}

// New creates new instance of actions
func New(userStorage storage.Stored, logger *zap.Logger) *processor {
	return &processor{
		storage: userStorage,
		logger:  logger,
	}
}
