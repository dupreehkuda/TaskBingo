//go:build unit

package pack_dislike_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_dislike"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/pack_dislike/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	s := mocks.NewMockstorage(ctrl)
	s.EXPECT().Dislike(gomock.Any(), "u1", "p1").Return(nil)
	require.NoError(t, pack_dislike.New(s, stubTx{}).Run(context.Background(), "u1", "p1"))
}
