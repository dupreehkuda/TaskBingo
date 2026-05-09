package app

import (
	"github.com/dupreehkuda/TaskBingo/pkg/passwordhash"
	"github.com/dupreehkuda/TaskBingo/pkg/tokens"
)

type tokenSigner struct {
	secret string
}

func newTokenSigner(secret string) *tokenSigner { return &tokenSigner{secret: secret} }

func (t *tokenSigner) Sign(userID, username string) (string, error) {
	return tokens.GenerateJWT(userID, username, t.secret)
}

type hasher struct{}

func (hasher) Hash(password, salt string) string { return passwordhash.Hash(password, salt) }
func (hasher) Salt(n int) (string, error)        { return passwordhash.Salt(n) }
