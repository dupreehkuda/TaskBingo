//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package user_login

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_login/storage"
)

type credsFetcher interface {
	FetchByUsername(ctx context.Context, username string) (storage.Credentials, error)
}

type hasher interface {
	Hash(password, salt string) string
}

type tokenSigner interface {
	Sign(userID, username string) (string, error)
}

type Usecase struct {
	storage credsFetcher
	hasher  hasher
	tokens  tokenSigner
	logger  *zap.Logger
}

func New(s credsFetcher, h hasher, t tokenSigner, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, hasher: h, tokens: t, logger: logger}
}

func (u *Usecase) Run(ctx context.Context, username, password string) (string, error) {
	creds, err := u.storage.FetchByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return "", errs.ErrWrongCredentials
		}
		u.logger.Error("FetchByUsername failed", zap.Error(err))
		return "", err
	}
	if u.hasher.Hash(password, creds.PasswordSalt) != creds.PasswordHash {
		return "", errs.ErrWrongCredentials
	}
	return u.tokens.Sign(creds.UserID, username)
}
