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
func (u userClient) LoginUser(userID, password string) (string, error) {
	data := api.LoginUserRequest{
		UserID:   &api.UUID{Id: userID},
		Password: password,
	}

	resp, err := u.conn.LoginUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)
	u.logger.Debug("incoming code", zap.Int("code", int(statusCode.Code())))

	if statusCode.Code() == codes.Unauthenticated {
		return "", errs.ErrWrongCredentials
	}

	return resp.UserID.Id, nil
}
