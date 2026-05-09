//go:build unit

package pack_get_many_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_many"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_many/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().GetMany(gomock.Any(), []string{"a", "b"}, "u1").Return(models.Packs{{ID: "a"}, {ID: "b"}}, nil)
	uc := pack_get_many.New(s, zap.NewNop())
	got, err := uc.Run(context.Background(), "u1", []string{"a", "b"})
	require.NoError(t, err)
	require.Len(t, got, 2)
}
