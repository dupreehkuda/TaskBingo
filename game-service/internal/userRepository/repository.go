package userRepository

import (
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
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
}

// userRepository provides connection to user service
type userRepository struct {
	conn   api.UsersClient
	logger *zap.Logger
}

// New returns an instance of userRepository
func New(address string, logger *zap.Logger) *userRepository {
	// Sleep before every container is awake
	time.Sleep(2 * time.Second)

	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Error when connecting to user-service (not connected)")
	}

	client := api.NewUsersClient(conn)

	return &userRepository{
		conn:   client,
		logger: logger,
	}
}
