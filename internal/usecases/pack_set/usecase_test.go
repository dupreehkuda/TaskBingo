//go:build unit

package pack_set_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_set"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_set/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Insert(gomock.Any(), gomock.Any(), "u1", "T", []string{"a", "b"}, true).Return(nil)

	uc := pack_set.New(s, stubTx{}, zap.NewNop())
	pack := &models.TaskPack{Pack: models.Pack{Title: "T", Tasks: []string{"a", "b"}}, IsPrivate: true}
	require.NoError(t, uc.Run(context.Background(), "u1", pack))
	require.NotEmpty(t, pack.ID)
}
