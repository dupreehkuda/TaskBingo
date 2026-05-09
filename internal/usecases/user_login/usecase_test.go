//go:build unit

package user_login_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_login"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_login/mocks"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_login/storage"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		setup   func(s *mocks.MockcredsFetcher, h *mocks.Mockhasher, ts *mocks.MocktokenSigner)
		wantErr error
	}{
		{
			name: "user not found -> wrong credentials",
			setup: func(s *mocks.MockcredsFetcher, _ *mocks.Mockhasher, _ *mocks.MocktokenSigner) {
				s.EXPECT().FetchByUsername(gomock.Any(), "alice").Return(storage.Credentials{}, errs.ErrNotFound)
			},
			wantErr: errs.ErrWrongCredentials,
		},
		{
			name: "wrong hash -> wrong credentials",
			setup: func(s *mocks.MockcredsFetcher, h *mocks.Mockhasher, _ *mocks.MocktokenSigner) {
				s.EXPECT().FetchByUsername(gomock.Any(), "alice").Return(storage.Credentials{UserID: "u1", PasswordHash: "H1", PasswordSalt: "S"}, nil)
				h.EXPECT().Hash("pw", "S").Return("H2")
			},
			wantErr: errs.ErrWrongCredentials,
		},
		{
			name: "ok -> returns token",
			setup: func(s *mocks.MockcredsFetcher, h *mocks.Mockhasher, ts *mocks.MocktokenSigner) {
				s.EXPECT().FetchByUsername(gomock.Any(), "alice").Return(storage.Credentials{UserID: "u1", PasswordHash: "H", PasswordSalt: "S"}, nil)
				h.EXPECT().Hash("pw", "S").Return("H")
				ts.EXPECT().Sign("u1", "alice").Return("jwt", nil)
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			s := mocks.NewMockcredsFetcher(ctrl)
			h := mocks.NewMockhasher(ctrl)
			ts := mocks.NewMocktokenSigner(ctrl)
			c.setup(s, h, ts)

			uc := user_login.New(s, h, ts, zap.NewNop())
			_, err := uc.Run(context.Background(), "alice", "pw")
			if c.wantErr != nil {
				require.ErrorIs(t, err, c.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
