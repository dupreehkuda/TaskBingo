//go:build unit

package user_get_all_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_all"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_all/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().ListAll(gomock.Any()).Return(models.Users{{UserID: "u1"}}, nil)

	uc := user_get_all.New(s, zap.NewNop())
	got, err := uc.Run(context.Background())
	require.NoError(t, err)
	require.Len(t, got, 1)
}
