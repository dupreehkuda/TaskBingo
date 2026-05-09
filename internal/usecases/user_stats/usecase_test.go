//go:build unit

package user_stats_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats/entity"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats/mocks"
)

func TestUsecase_Run(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	const me = "u1"

	now := time.Date(2026, 5, 9, 12, 0, 0, 0, time.UTC)
	row := func(kind string, when time.Time, bingo int, nums []int, packID, packTitle, winner, otherUser string) entity.GameRow {
		return entity.GameRow{
			GameID:       "g-" + when.Format("0102") + kind,
			Kind:         kind,
			Finished:     when,
			PackID:       packID,
			PackTitle:    packTitle,
			Winner:       winner,
			User1ID:      me,
			User2ID:      otherUser,
			User1Bingo:   bingo,
			User2Bingo:   0,
			User1Numbers: nums,
			User2Numbers: []int{},
		}
	}

	t.Run("aggregates 7d window into 7 daily buckets", func(t *testing.T) {
		s := mocks.NewMockstorage(ctrl)
		rows := []entity.GameRow{
			row("solo", now.Add(-6*24*time.Hour), 2, []int{1, 2, 3, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "p1", "Pack One", "", ""),
			row("duo", now.Add(-1*24*time.Hour), 1, []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "p1", "Pack One", me, "u2"),
			row("duo", now.Add(-30*time.Minute), 3, []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, "p2", "Pack Two", "u2", "u2"),
		}
		s.EXPECT().FetchFinishedGames(gomock.Any(), me, gomock.Any()).Return(rows, nil)

		uc := user_stats.New(s, zap.NewNop())
		uc.SetClock(func() time.Time { return now })

		got, err := uc.Run(context.Background(), me, 7)
		require.NoError(t, err)
		require.Len(t, got.Buckets, 7)

		// summary
		require.Equal(t, 3, got.Summary.Games)
		require.Equal(t, 1, got.Summary.Wins)
		require.Equal(t, 1, got.Summary.Losses)
		require.Equal(t, 6, got.Summary.TotalBingo)
		require.Equal(t, 2, got.Summary.SoloBingo)
		require.Equal(t, 4, got.Summary.DuoBingo)
		// 4 + 5 + 8 = 17 closed tasks
		require.Equal(t, 17, got.Summary.TasksClosed)

		// top packs: p1 played twice, p2 once
		require.Len(t, got.TopPacks, 2)
		require.Equal(t, "p1", got.TopPacks[0].ID)
		require.Equal(t, 2, got.TopPacks[0].Count)
		require.Equal(t, "p2", got.TopPacks[1].ID)
		require.Equal(t, 1, got.TopPacks[1].Count)

		// most recent bucket holds today's duo win
		last := got.Buckets[6]
		require.Equal(t, 3, last.DuoBingo)
		require.Equal(t, 8, last.TasksClosed)
	})

	t.Run("30d window yields 5 weekly buckets", func(t *testing.T) {
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().FetchFinishedGames(gomock.Any(), me, gomock.Any()).Return([]entity.GameRow{}, nil)

		uc := user_stats.New(s, zap.NewNop())
		uc.SetClock(func() time.Time { return now })

		got, err := uc.Run(context.Background(), me, 30)
		require.NoError(t, err)
		require.Len(t, got.Buckets, 5)
		require.Equal(t, 30, got.Window.Days)
	})

	t.Run("propagates storage error", func(t *testing.T) {
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().FetchFinishedGames(gomock.Any(), me, gomock.Any()).Return(nil, errors.New("boom"))
		uc := user_stats.New(s, zap.NewNop())
		uc.SetClock(func() time.Time { return now })
		_, err := uc.Run(context.Background(), me, 7)
		require.Error(t, err)
	})

	t.Run("days param defaults and clamps", func(t *testing.T) {
		s := mocks.NewMockstorage(ctrl)
		s.EXPECT().FetchFinishedGames(gomock.Any(), me, gomock.Any()).Return([]entity.GameRow{}, nil).Times(3)
		uc := user_stats.New(s, zap.NewNop())
		uc.SetClock(func() time.Time { return now })

		zero, _ := uc.Run(context.Background(), me, 0)
		require.Equal(t, 7, zero.Window.Days)

		neg, _ := uc.Run(context.Background(), me, -5)
		require.Equal(t, 7, neg.Window.Days)

		huge, _ := uc.Run(context.Background(), me, 1000)
		require.Equal(t, 30, huge.Window.Days)
	})
}
