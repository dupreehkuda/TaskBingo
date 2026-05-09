//go:build unit

package comment_add_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_add"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_add/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Insert(gomock.Any(), gomock.AssignableToTypeOf(&models.Comment{})).
		DoAndReturn(func(_ context.Context, c *models.Comment) (*models.Comment, error) {
			require.NotEmpty(t, c.ID)
			require.Equal(t, "g1", c.GameID)
			require.Equal(t, "u1", c.UserID)
			require.Equal(t, "hello", c.Body)
			return c, nil
		})
	got, err := comment_add.New(s, zap.NewNop()).Run(context.Background(), "u1", "g1", "hello")
	require.NoError(t, err)
	require.Equal(t, "hello", got.Body)
}
