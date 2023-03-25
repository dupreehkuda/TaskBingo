package userClient

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// LoginUser sends login request to user service
func (u userClient) LoginUser(username, password string) (string, error) {
	data := api.LoginUserRequest{
		Username: username,
		Password: password,
	}

	resp, err := u.conn.LoginUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)

	if statusCode.Code() == codes.Unauthenticated {
		return "", errs.ErrWrongCredentials
	}

	return resp.UserID, nil
}
