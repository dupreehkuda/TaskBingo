//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package user_register

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type storage interface {
	CheckDuplicate(ctx context.Context, username, email string) (bool, error)
	Insert(ctx context.Context, userID, username, email, city, passwordHash, passwordSalt string) error
}

type hasher interface {
	Hash(password, salt string) string
	Salt(n int) (string, error)
}

type tokenSigner interface {
	Sign(userID, username string) (string, error)
}

type txManager interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Usecase struct {
	storage storage
	hasher  hasher
	tokens  tokenSigner
	tx      txManager
	logger  *zap.Logger
}

func New(s storage, h hasher, t tokenSigner, tx txManager, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, hasher: h, tokens: t, tx: tx, logger: logger}
}

// Run registers a new user and returns a freshly minted JWT.
func (u *Usecase) Run(ctx context.Context, c *models.RegisterCredentials) (string, error) {
	exists, err := u.storage.CheckDuplicate(ctx, c.Username, c.Email)
	if err != nil {
		u.logger.Error("CheckDuplicate failed", zap.Error(err))
		return "", err
	}
	if exists {
		return "", errs.ErrCredentialsInUse
	}

	salt, err := u.hasher.Salt(10)
	if err != nil {
		return "", err
	}
	hash := u.hasher.Hash(c.Password, salt)

	userID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	city := cases.Title(language.English).String(c.City)

	if err := u.tx.RunInTx(ctx, func(ctx context.Context) error {
		return u.storage.Insert(ctx, userID.String(), c.Username, c.Email, city, hash, salt)
	}); err != nil {
		u.logger.Error("Insert user failed", zap.Error(err))
		return "", err
	}

	return u.tokens.Sign(userID.String(), c.Username)
}
