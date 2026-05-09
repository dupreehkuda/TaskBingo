//go:build unit

package user_register_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_register"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_register/mocks"
)

type stubTx struct{}

func (stubTx) RunInTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	creds := &models.RegisterCredentials{Username: "alice", Email: "a@x.io", City: "moscow", Password: "pw"}

	tests := []struct {
		name    string
		setup   func(s *mocks.Mockstorage, h *mocks.Mockhasher, ts *mocks.MocktokenSigner)
		wantErr error
	}{
		{
			name: "duplicate -> ErrCredentialsInUse",
			setup: func(s *mocks.Mockstorage, _ *mocks.Mockhasher, _ *mocks.MocktokenSigner) {
				s.EXPECT().CheckDuplicate(gomock.Any(), "alice", "a@x.io").Return(true, nil)
			},
			wantErr: errs.ErrCredentialsInUse,
		},
		{
			name: "happy path",
			setup: func(s *mocks.Mockstorage, h *mocks.Mockhasher, ts *mocks.MocktokenSigner) {
				s.EXPECT().CheckDuplicate(gomock.Any(), "alice", "a@x.io").Return(false, nil)
				h.EXPECT().Salt(10).Return("SALT123456", nil)
				h.EXPECT().Hash("pw", "SALT123456").Return("HASHED")
				s.EXPECT().Insert(gomock.Any(), gomock.Any(), "alice", "a@x.io", "Moscow", "HASHED", "SALT123456").Return(nil)
				ts.EXPECT().Sign(gomock.Any(), "alice").Return("jwt-token", nil)
			},
		},
		{
			name: "duplicate-check storage error propagates",
			setup: func(s *mocks.Mockstorage, _ *mocks.Mockhasher, _ *mocks.MocktokenSigner) {
				s.EXPECT().CheckDuplicate(gomock.Any(), "alice", "a@x.io").Return(false, errors.New("boom"))
			},
			wantErr: errors.New("boom"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			s := mocks.NewMockstorage(ctrl)
			h := mocks.NewMockhasher(ctrl)
			ts := mocks.NewMocktokenSigner(ctrl)
			tt.setup(s, h, ts)

			uc := user_register.New(s, h, ts, stubTx{}, zap.NewNop())
			_, err := uc.Run(context.Background(), creds)
			if tt.wantErr != nil {
				require.Error(t, err)
				switch {
				case errors.Is(tt.wantErr, errs.ErrCredentialsInUse):
					require.ErrorIs(t, err, errs.ErrCredentialsInUse)
				default:
					require.ErrorContains(t, err, tt.wantErr.Error())
				}
				return
			}
			require.NoError(t, err)
		})
	}
}
