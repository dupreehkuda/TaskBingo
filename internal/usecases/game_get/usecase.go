//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package game_get

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	Get(ctx context.Context, gameID string) (*models.Game, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, l *zap.Logger) *Usecase { return &Usecase{storage: s, logger: l} }

func (u *Usecase) Run(ctx context.Context, gameID string) (*models.Game, error) {
	return u.storage.Get(ctx, gameID)
}
