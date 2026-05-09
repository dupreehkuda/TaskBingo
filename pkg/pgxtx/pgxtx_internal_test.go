package pgxtx

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

// fakeTx is a no-op pgx.Tx used purely as a sentinel value for context testing.
type fakeTx struct{ pgx.Tx }

func TestRunInTx_NestedReusesExistingTx(t *testing.T) {
	var ft fakeTx
	ctx := context.WithValue(context.Background(), ctxKey{}, pgx.Tx(&ft))

	m := &Manager{pool: nil} // pool unused on the nested path
	sentinel := errors.New("from-fn")
	err := m.RunInTx(ctx, func(c context.Context) error {
		got, ok := From(c)
		require.True(t, ok)
		require.Equal(t, pgx.Tx(&ft), got)
		return sentinel
	})
	require.ErrorIs(t, err, sentinel)
}
