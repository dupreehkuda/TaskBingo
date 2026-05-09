//go:build unit

package game_realtime_action_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_realtime_action"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_realtime_action/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run_NoBroadcastOnEmptyRoom(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	uc := game_realtime_action.New(s, stubTx{}, zap.NewNop())
	r := &models.Room{Status: models.GameCreated}
	update, err := uc.Run(context.Background(), r, &models.GameAction{})
	require.NoError(t, err)
	require.Nil(t, update)
}

func TestUsecase_LoadRoom(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Get(gomock.Any(), "g1").Return(&models.Game{GameID: "g1"}, nil)

	uc := game_realtime_action.New(s, stubTx{}, zap.NewNop())
	r, err := uc.LoadRoom(context.Background(), "g1")
	require.NoError(t, err)
	require.Equal(t, "g1", r.Id)
	require.Equal(t, models.GameCreated, r.Status)
}
