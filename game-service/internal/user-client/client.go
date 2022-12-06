package user_client

import (
	"context"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type userClient struct {
	conn   api.UsersClient
	logger *zap.Logger
}

// New returns an instance of userClient
func New(logger *zap.Logger) *userClient {
	// Sleep before every container is awake
	time.Sleep(3 * time.Second)

	creds := insecure.NewCredentials()
	conn, err := grpc.Dial("user-service:8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Error when connecting to user-service (not connected)")
	}

	client := api.NewUsersClient(conn)

	return &userClient{
		conn:   client,
		logger: logger,
	}
}

func (u userClient) GetUser(userId string) (*models.Response, error) {
	resp, err := u.conn.Get(context.Background(), &api.GetRequest{Id: userId})
	if err != nil {
		u.logger.Error("Error when getting user data")
		return nil, err
	}

	return &models.Response{
		UserID: resp.Nickname,
		Points: int(resp.Points),
		Email:  resp.Email,
	}, nil
}
