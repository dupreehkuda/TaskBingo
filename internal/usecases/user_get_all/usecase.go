//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package user_get_all

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	ListAll(ctx context.Context) (models.Users, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, l *zap.Logger) *Usecase { return &Usecase{storage: s, logger: l} }

func (u *Usecase) Run(ctx context.Context) (models.Users, error) {
	return u.storage.ListAll(ctx)
}
