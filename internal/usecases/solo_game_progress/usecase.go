//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package solo_game_progress

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type storage interface {
	Update(ctx context.Context, userID, gameID string, userNumbers []int32) (pgconn.CommandTag, error)
}

type Usecase struct {
	storage storage
}

func New(s storage) *Usecase { return &Usecase{storage: s} }

func (u *Usecase) Run(ctx context.Context, userID, gameID string, userNumbers []int32) error {
	_, err := u.storage.Update(ctx, userID, gameID, userNumbers)
	return err
}
