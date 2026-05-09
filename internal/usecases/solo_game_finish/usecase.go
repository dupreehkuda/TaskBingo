//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package solo_game_finish

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_finish/entity"
)

type storage interface {
	Finalize(ctx context.Context, userID, gameID string, userNumbers []int32, bingo int32) (int64, error)
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

// Run counts the bingo on the submitted user numbers and persists the final
// state in a tx (game UPDATE + users.solo_bingo bump). Returns ErrNotFound
// if the game doesn't belong to the user or isn't a solo game.
func (u *Usecase) Run(ctx context.Context, userID, gameID string, userNumbers []int32) (*models.SoloFinishResponse, error) {
	bingo := entity.CountBingo(userNumbers)

	var rows int64
	if err := u.tx.RunInTx(ctx, func(ctx context.Context) error {
		var err error
		rows, err = u.storage.Finalize(ctx, userID, gameID, userNumbers, bingo)
		return err
	}); err != nil {
		u.logger.Error("solo finalize failed", zap.Error(err))
		return nil, err
	}
	if rows == 0 {
		return nil, errs.ErrNotFound
	}
	return &models.SoloFinishResponse{Bingo: bingo}, nil
}
