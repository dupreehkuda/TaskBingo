package userClient

import (
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// UserDataClient is an interface for user data service
type UserDataClient interface {
	GetUserData(login string) (*models.UserAccountInfoResponse, error)
	RegisterUser(creds *models.RegisterCredentials) error
	LoginUser(login, password string) error
	GetRatedPacks() ([]string, error)
	LikeTaskPack(login, pack string) error
	DislikeTaskPack(login, pack string) error
	RateTaskPack(login, pack string) error
	UnrateTaskPack(login, pack string) error
	AssignNewPack(login, pack string) error
	GetAllUsers() (*[]models.User, error)
	AcceptFriend(login, person string) error
	DeleteFriend(login, person string) error
	RequestFriend(login, person string) error
	CreateGame(game *models.Game) error
}

// userClient provides connection to user service
type userClient struct {
	conn   api.UsersClient
	logger *zap.Logger
}

// New returns an instance of userClient
func New(address string, logger *zap.Logger) *userClient {
	// Sleep before every container is awake
	time.Sleep(2 * time.Second)

	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Error when connecting to user-service (not connected)")
	}

	client := api.NewUsersClient(conn)

	return &userClient{
		conn:   client,
		logger: logger,
	}
}
