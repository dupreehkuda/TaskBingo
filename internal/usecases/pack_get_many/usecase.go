//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package pack_get_many

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	GetMany(ctx context.Context, packIDs []string, requesterID string) (models.Packs, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, l *zap.Logger) *Usecase { return &Usecase{storage: s, logger: l} }

func (u *Usecase) Run(ctx context.Context, requesterID string, packIDs []string) (models.Packs, error) {
	return u.storage.GetMany(ctx, packIDs, requesterID)
}
