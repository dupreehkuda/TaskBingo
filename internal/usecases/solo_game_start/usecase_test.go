//go:build unit

package solo_game_start_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_start"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_start/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&models.Game{})).Return(nil)

	resp, err := solo_game_start.New(s, stubTx{}, zap.NewNop()).
		Run(context.Background(), "u1", "p1")
	require.NoError(t, err)
	require.NotEmpty(t, resp.GameID)
	require.Len(t, resp.Numbers, 16)
}
