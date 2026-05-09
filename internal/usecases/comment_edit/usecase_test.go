//go:build unit

package comment_edit_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_edit"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_edit/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()

	t.Run("ok returns updated row", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().Update(gomock.Any(), "c1", "u1", "new").Return(&models.Comment{ID: "c1", Body: "new"}, nil)
		got, err := comment_edit.New(s, zap.NewNop()).Run(context.Background(), "u1", "c1", "new")
		require.NoError(t, err)
		require.Equal(t, "new", got.Body)
	})

	t.Run("ErrNoRows -> ErrNotFound", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().Update(gomock.Any(), "c1", "u1", "x").Return(nil, pgx.ErrNoRows)
		_, err := comment_edit.New(s, zap.NewNop()).Run(context.Background(), "u1", "c1", "x")
		require.ErrorIs(t, err, errs.ErrNotFound)
	})
}
