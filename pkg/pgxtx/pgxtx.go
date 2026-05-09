package pgxtx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ctxKey struct{}

// From returns the pgx.Tx attached to ctx, if any.
func From(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(ctxKey{}).(pgx.Tx)
	return tx, ok
}

// Manager owns a pool and starts transactions.
type Manager struct {
	pool *pgxpool.Pool
}

// New returns a Manager bound to the given pool.
func New(pool *pgxpool.Pool) *Manager {
	return &Manager{pool: pool}
}

// RunInTx begins a tx, attaches it to ctx, and runs fn. Commits on nil error,
// rolls back otherwise. Nested calls reuse the existing tx.
func (m *Manager) RunInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	if _, ok := From(ctx); ok {
		return fn(ctx)
	}
	tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	txCtx := context.WithValue(ctx, ctxKey{}, tx)
	if err := fn(txCtx); err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}
	return tx.Commit(ctx)
}
