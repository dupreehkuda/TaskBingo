package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/storage"
)

// Processor is interface for business-logic
type Processor interface {
	RegisterUser(login, password, email, city string) error
	LoginUser(login, password string) error
	GetUserData(userId string) (*models.GetUserDataResponse, error)
	GetRatedPacks() ([]string, error)
	LikePack(login, pack string, inc int) error
	RatePack(login, pack string, inc int) error
	AssignNewPack(login, pack string) error
	GetAllUsers() (*[]models.AllUsers, error)
	AcceptFriend(login, person string) error
	DeleteFriend(login, person string) error
	RequestFriend(login, person string) error
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
