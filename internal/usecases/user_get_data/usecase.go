//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package user_get_data

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	Get(ctx context.Context, userID string) (*models.UserAccountInfo, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, logger: logger}
}

func (u *Usecase) Run(ctx context.Context, userID string) (*models.UserAccountInfo, error) {
	return u.storage.Get(ctx, userID)
}
