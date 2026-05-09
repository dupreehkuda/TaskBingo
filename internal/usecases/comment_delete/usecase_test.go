//go:build unit

package comment_delete_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_delete"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_delete/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().Delete(gomock.Any(), "c1", "u1").Return(pgconn.NewCommandTag("DELETE 1"), nil)
		require.NoError(t, comment_delete.New(s).Run(context.Background(), "u1", "c1"))
	})

	t.Run("zero rows -> ErrNotFound", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().Delete(gomock.Any(), "c1", "u1").Return(pgconn.NewCommandTag("DELETE 0"), nil)
		require.ErrorIs(t, comment_delete.New(s).Run(context.Background(), "u1", "c1"), errs.ErrNotFound)
	})
}
