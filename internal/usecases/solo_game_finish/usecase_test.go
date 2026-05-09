//go:build unit

package solo_game_finish_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_finish"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_finish/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()

	t.Run("happy path counts bingo and finalises", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		s := mocks.NewMockstorage(ctrl)
		// One full top row -> bingo == 1.
		nums := []int32{1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		s.EXPECT().Finalize(gomock.Any(), "u1", "g1", nums, int32(1)).Return(int64(1), nil)

		resp, err := solo_game_finish.New(s, stubTx{}, zap.NewNop()).
			Run(context.Background(), "u1", "g1", nums)
		require.NoError(t, err)
		require.Equal(t, int32(1), resp.Bingo)
	})

	t.Run("zero rows -> ErrNotFound", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		s := mocks.NewMockstorage(ctrl)
		nums := make([]int32, 16)
		s.EXPECT().Finalize(gomock.Any(), "u1", "g1", nums, int32(0)).Return(int64(0), nil)
		_, err := solo_game_finish.New(s, stubTx{}, zap.NewNop()).
			Run(context.Background(), "u1", "g1", nums)
		require.ErrorIs(t, err, errs.ErrNotFound)
	})
}
