package repository

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
)

// LoginUser sends login request to user service
func (r repository) LoginUser(ctx context.Context, username, password string) (string, error) {
	resp, err := r.conn.LoginUser(ctx, mapToLogin(username, password))

	statusCode, _ := status.FromError(err)

	if statusCode.Code() == codes.Unauthenticated {
		return "", errs.ErrWrongCredentials
	}

	return resp.UserID, nil
}
