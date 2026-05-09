//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package pack_set

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type storage interface {
	Insert(ctx context.Context, packID, creatorID, title string, tasks []string, isPrivate bool) error
}

type txManager interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Usecase struct {
	storage storage
	tx      txManager
	logger  *zap.Logger
}

func New(s storage, tx txManager, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, tx: tx, logger: logger}
}

func (u *Usecase) Run(ctx context.Context, userID string, pack *models.TaskPack) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	pack.ID = id.String()
	return u.tx.RunInTx(ctx, func(ctx context.Context) error {
		return u.storage.Insert(ctx, pack.ID, userID, pack.Pack.Title, pack.Pack.Tasks, pack.IsPrivate)
	})
}
