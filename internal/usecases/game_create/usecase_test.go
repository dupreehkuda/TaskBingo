//go:build unit

package game_create_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_create"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_create/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&models.Game{})).Return(nil)

	short, err := game_create.New(s, stubTx{}, zap.NewNop()).
		Run(context.Background(), "u1", "u2", "p1")
	require.NoError(t, err)
	require.NotEmpty(t, short.GameID)
	require.Equal(t, "u1", short.User1Id)
	require.Equal(t, "u2", short.User2Id)
}
