//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package game_accept

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type storage interface {
	Accept(ctx context.Context, userID, gameID string) (pgconn.CommandTag, error)
}

type Usecase struct {
	storage storage
}

func New(s storage) *Usecase { return &Usecase{storage: s} }

func (u *Usecase) Run(ctx context.Context, userID, gameID string) error {
	_, err := u.storage.Accept(ctx, userID, gameID)
	return err
}
