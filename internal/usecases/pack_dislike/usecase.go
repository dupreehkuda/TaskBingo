//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package pack_dislike

import "context"

type storage interface {
	Dislike(ctx context.Context, userID, packID string) error
}

type txManager interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Usecase struct {
	storage storage
	tx      txManager
}

func New(s storage, tx txManager) *Usecase { return &Usecase{storage: s, tx: tx} }

func (u *Usecase) Run(ctx context.Context, userID, packID string) error {
	return u.tx.RunInTx(ctx, func(ctx context.Context) error {
		return u.storage.Dislike(ctx, userID, packID)
	})
}
