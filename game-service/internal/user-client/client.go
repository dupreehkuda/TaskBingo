package user_client

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

type userClient struct {
	conn   api.UsersClient
	logger *zap.Logger
}

// New returns an instance of userClient
func New(address string, logger *zap.Logger) *userClient {
	// Sleep before every container is awake
	time.Sleep(3 * time.Second)

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

func (u userClient) GetUserData(login string) (*models.Response, error) {
	resp, err := u.conn.GetUserData(context.Background(), &api.GetUserDataRequest{Login: login})
	if err != nil {
		u.logger.Error("Error when getting user data")
		return nil, err
	}

	return &models.Response{
		UserID: resp.Login,
		Points: int(resp.Points),
		Email:  resp.Email,
	}, nil
}

func (u userClient) RegisterUser(login, email, password string) error {
	data := api.RegisterUserRequest{
		Login:    login,
		Email:    email,
		Password: password,
	}

	_, err := u.conn.RegisterUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)

	if statusCode.Code() == codes.AlreadyExists {
		return errs.ErrCredentialsInUse
	}

	return nil
}

func (u userClient) LoginUser(login, password string) error {
	data := api.LoginUserRequest{
		Login:    login,
		Password: password,
	}

	_, err := u.conn.LoginUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)

	if statusCode.Code() == codes.Unauthenticated {
		return errs.ErrWrongCredentials
	}

	return nil
}
