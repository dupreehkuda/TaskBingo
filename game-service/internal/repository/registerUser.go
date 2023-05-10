package repository

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// RegisterUser sends register request to user service
func (r repository) RegisterUser(ctx context.Context, credits *models.RegisterCredentials) (string, error) {
	resp, err := r.conn.RegisterUser(ctx, mapToRegister(credits))

	statusCode, _ := status.FromError(err)
	r.logger.Debug("incoming code", zap.Int("code", int(statusCode.Code())))

	if statusCode.Code() == codes.AlreadyExists {
		return "", errs.ErrCredentialsInUse
	}

	return resp.UserID, nil
}
