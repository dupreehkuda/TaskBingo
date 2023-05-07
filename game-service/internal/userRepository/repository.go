package userRepository

import (
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

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
