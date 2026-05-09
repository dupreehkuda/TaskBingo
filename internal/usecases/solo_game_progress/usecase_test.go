//go:build unit

package solo_game_progress_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_progress"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_progress/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Update(gomock.Any(), "u1", "g1", []int32{1, 0, 0}).Return(pgconn.CommandTag{}, nil)
	require.NoError(t, solo_game_progress.New(s).Run(context.Background(), "u1", "g1", []int32{1, 0, 0}))
}
