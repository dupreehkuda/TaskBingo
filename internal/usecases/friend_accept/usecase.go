//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package friend_accept

import "context"

type storage interface {
	Accept(ctx context.Context, userID, friendID string) error
}

type txManager interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Usecase struct {
	storage storage
	tx      txManager
}

func New(s storage, tx txManager) *Usecase { return &Usecase{storage: s, tx: tx} }

func (u *Usecase) Run(ctx context.Context, userID, friendID string) error {
	return u.tx.RunInTx(ctx, func(ctx context.Context) error {
		return u.storage.Accept(ctx, userID, friendID)
	})
}
