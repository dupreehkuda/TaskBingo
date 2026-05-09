//go:build unit

package user_get_data_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_data"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_get_data/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)

	t.Run("propagates response", func(t *testing.T) {
		want := &models.UserAccountInfo{UserID: "u1"}
		s.EXPECT().Get(gomock.Any(), "u1").Return(want, nil)
		uc := user_get_data.New(s, zap.NewNop())
		got, err := uc.Run(context.Background(), "u1")
		require.NoError(t, err)
		require.Equal(t, want, got)
	})

	t.Run("propagates error", func(t *testing.T) {
		s.EXPECT().Get(gomock.Any(), "u2").Return(nil, errors.New("boom"))
		uc := user_get_data.New(s, zap.NewNop())
		_, err := uc.Run(context.Background(), "u2")
		require.Error(t, err)
	})
}
