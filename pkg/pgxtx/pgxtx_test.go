package pgxtx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dupreehkuda/TaskBingo/pkg/pgxtx"
)

func TestFrom_NoTx(t *testing.T) {
	_, ok := pgxtx.From(context.Background())
	require.False(t, ok)
}
