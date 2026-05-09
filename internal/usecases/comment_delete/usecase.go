//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package comment_delete

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type storage interface {
	Delete(ctx context.Context, commentID, userID string) (pgconn.CommandTag, error)
}

type Usecase struct {
	storage storage
}

func New(s storage) *Usecase { return &Usecase{storage: s} }

func (u *Usecase) Run(ctx context.Context, userID, commentID string) error {
	tag, err := u.storage.Delete(ctx, commentID, userID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errs.ErrNotFound
	}
	return nil
}
