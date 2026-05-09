//go:build unit

package pack_get_rated_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_rated"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_get_rated/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().TopRated(gomock.Any()).Return(models.Packs{{ID: "p"}}, nil)
	uc := pack_get_rated.New(s, zap.NewNop())
	got, err := uc.Run(context.Background())
	require.NoError(t, err)
	require.Len(t, got, 1)
}
