//go:build unit

package game_delete_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_delete"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_delete/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Delete(gomock.Any(), "u1", "g1").Return(pgconn.CommandTag{}, nil)
	require.NoError(t, game_delete.New(s).Run(context.Background(), "u1", "g1"))
}
