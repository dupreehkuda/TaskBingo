package uuidcheck_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

func TestCheck(t *testing.T) {
	t.Run("empty rejects", func(t *testing.T) {
		require.ErrorIs(t, uuidcheck.Check(""), errs.ErrEmptyRequest)
	})
	t.Run("malformed rejects", func(t *testing.T) {
		require.Error(t, uuidcheck.Check("not-a-uuid"))
	})
	t.Run("valid passes", func(t *testing.T) {
		require.NoError(t, uuidcheck.Check("c0a801c1-1f3b-4a0b-9d7c-12d34f5e6a7b"))
	})
	t.Run("multi rejects on first bad", func(t *testing.T) {
		require.Error(t, uuidcheck.Check("c0a801c1-1f3b-4a0b-9d7c-12d34f5e6a7b", "x"))
	})
}
