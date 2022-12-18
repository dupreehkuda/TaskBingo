package userClient

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// LoginUser sends login request to user service
func (u userClient) LoginUser(login, password string) error {
	data := api.LoginUserRequest{
		Login:    login,
		Password: password,
	}

	_, err := u.conn.LoginUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)
	u.logger.Debug("incoming code", zap.Int("code", int(statusCode.Code())))

	if statusCode.Code() == codes.Unauthenticated {
		return errs.ErrWrongCredentials
	}

	return nil
}
