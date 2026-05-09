//go:build unit

package comment_list_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_list"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/comment_list/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().List(gomock.Any(), "g1", "u1").Return([]models.Comment{{ID: "c1"}}, nil)

	got, err := comment_list.New(s, zap.NewNop()).Run(context.Background(), "u1", "g1")
	require.NoError(t, err)
	require.Len(t, got, 1)
}
