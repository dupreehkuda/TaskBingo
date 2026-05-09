//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package comment_add

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	Insert(ctx context.Context, c *models.Comment) (*models.Comment, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, logger: logger}
}

func (u *Usecase) Run(ctx context.Context, userID, gameID, body string) (*models.Comment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	c := &models.Comment{
		ID:     id.String(),
		GameID: gameID,
		UserID: userID,
		Body:   body,
	}
	out, err := u.storage.Insert(ctx, c)
	if err != nil {
		u.logger.Error("comment insert failed", zap.Error(err))
		return nil, err
	}
	return out, nil
}
