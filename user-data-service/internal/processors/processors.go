package processors

import (
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/storage"
)

// Processor is interface for business-logic
type Processor interface {
	RegisterUser(username, password, email, city string) (string, error)
	LoginUser(username, password string) (string, error)
	GetUserData(userID uuid.UUID) (*models.GetUserDataResponse, error)
	GetRatedPacks() ([]uuid.UUID, error)
	LikePack(userID, pack uuid.UUID, inc int) error
	RatePack(userID, pack uuid.UUID, inc int) error
	AssignNewPack(userID, packID uuid.UUID, packName string) error
	GetAllUsers() (*[]models.AllUsers, error)
	AcceptFriend(userID, friendID uuid.UUID) error
	DeleteFriend(userID, friendID uuid.UUID) error
	RequestFriend(userID, friendID uuid.UUID) error
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
