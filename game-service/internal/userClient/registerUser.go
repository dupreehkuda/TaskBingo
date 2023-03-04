package userClient

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// RegisterUser sends register request to user service
func (u userClient) RegisterUser(credits *models.RegisterCredentials) error {
	data := api.RegisterUserRequest{
		UserID:   &api.UUID{Id: credits.UserID},
		Username: credits.Username,
		Email:    credits.Email,
		City:     credits.City,
		Password: credits.Password,
	}

	_, err := u.conn.RegisterUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)
	u.logger.Debug("incoming code", zap.Int("code", int(statusCode.Code())))

	if statusCode.Code() == codes.AlreadyExists {
		return errs.ErrCredentialsInUse
	}

	return nil
}
