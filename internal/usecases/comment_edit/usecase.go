//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package comment_edit

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type storage interface {
	Update(ctx context.Context, commentID, userID, body string) (*models.Comment, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
}

func New(s storage, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, logger: logger}
}

func (u *Usecase) Run(ctx context.Context, userID, commentID, body string) (*models.Comment, error) {
	out, err := u.storage.Update(ctx, commentID, userID, body)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		u.logger.Error("comment update failed", zap.Error(err))
		return nil, err
	}
	return out, nil
}
