//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package comment_list

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	List(ctx context.Context, gameID, userID string) ([]models.Comment, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, logger: logger}
}

func (u *Usecase) Run(ctx context.Context, userID, gameID string) ([]models.Comment, error) {
	return u.storage.List(ctx, gameID, userID)
}
