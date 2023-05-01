package userRepository

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
func (u userRepository) RegisterUser(credits *models.RegisterCredentials) (string, error) {
	data := api.RegisterUserRequest{
		Username: credits.Username,
		Email:    credits.Email,
		City:     credits.City,
		Password: credits.Password,
	}

	resp, err := u.conn.RegisterUser(context.Background(), &data)

	statusCode, _ := status.FromError(err)
	u.logger.Debug("incoming code", zap.Int("code", int(statusCode.Code())))

	if statusCode.Code() == codes.AlreadyExists {
		return "", errs.ErrCredentialsInUse
	}

	return resp.UserID, nil
}
