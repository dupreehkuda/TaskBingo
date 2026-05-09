//go:build unit

package friend_delete_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/friend_delete"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/friend_delete/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Delete(gomock.Any(), "u1", "u2").Return(nil)
	require.NoError(t, friend_delete.New(s, stubTx{}).Run(context.Background(), "u1", "u2"))
}
