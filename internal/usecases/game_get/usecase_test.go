//go:build unit

package game_get_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_get"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_get/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	want := &models.Game{GameID: "g1"}
	s.EXPECT().Get(gomock.Any(), "g1").Return(want, nil)
	got, err := game_get.New(s, zap.NewNop()).Run(context.Background(), "g1")
	require.NoError(t, err)
	require.Equal(t, want, got)
}
