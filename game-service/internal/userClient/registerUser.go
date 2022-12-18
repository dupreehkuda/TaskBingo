package userClient

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (u userClient) RegisterUser(login, email, password string) error {
	data := api.RegisterUserRequest{
		Login:    login,
		Email:    email,
		Password: password,
	}

	_, err := u.conn.RegisterUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)
	u.logger.Debug("incoming code", zap.Int("code", int(statusCode.Code())))

	if statusCode.Code() == codes.AlreadyExists {
		return errs.ErrCredentialsInUse
	}

	return nil
}
