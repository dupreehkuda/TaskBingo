//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package user_stats

import (
	"context"
	"sort"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats/entity"
)

type storage interface {
	FetchFinishedGames(ctx context.Context, userID string, from time.Time) ([]entity.GameRow, error)
}

type Usecase struct {
	storage storage
	logger  *zap.Logger
	now     func() time.Time
}

func New(s storage, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, logger: logger, now: time.Now}
}

// SetClock overrides the clock used to compute the window. Tests use this to
// pin "now" so bucket assertions are deterministic.
func (u *Usecase) SetClock(now func() time.Time) {
	u.now = now
}

const (
	bucketCount7d  = 7
	bucketCount30d = 5
	defaultDays    = 7
	maxDays        = 30
	topPackLimit   = 5
)

// Run aggregates the user's finished games over the requested window into
// buckets, a top-packs list, and a summary tile row.
//
//nolint:gocyclo // single sequential aggregation
func (u *Usecase) Run(ctx context.Context, userID string, days int) (*entity.Result, error) {
	days = clampDays(days)
	to := u.now().UTC()
	bucketN := bucketCount(days)
	bucketSpan := time.Duration(days) / time.Duration(bucketN) * 24 * time.Hour
	from := to.Add(-time.Duration(days) * 24 * time.Hour)

	rows, err := u.storage.FetchFinishedGames(ctx, userID, from)
	if err != nil {
		return nil, err
	}

	buckets := make([]entity.Bucket, bucketN)
	for i := range buckets {
		buckets[i].Label = bucketLabel(to, days, bucketN, i)
	}

	packCounts := make(map[string]*entity.PackCount, 8)
	var summary entity.Summary

	for _, g := range rows {
		mine := pickSide(g, userID)
		idx := bucketIndex(g.Finished, to, bucketSpan, bucketN)
		if idx < 0 || idx >= bucketN {
			continue
		}

		tasks := countNonZero(mine.numbers)
		buckets[idx].TasksClosed += tasks
		summary.TasksClosed += tasks
		summary.TotalBingo += mine.bingo

		switch g.Kind {
		case "solo":
			buckets[idx].SoloBingo += mine.bingo
			summary.SoloBingo += mine.bingo
		default:
			buckets[idx].DuoBingo += mine.bingo
			summary.DuoBingo += mine.bingo
		}

		summary.Games++
		if g.Kind != "solo" {
			if g.Winner == userID {
				summary.Wins++
			} else if g.Winner != "" {
				summary.Losses++
			}
		}

		if pc, ok := packCounts[g.PackID]; ok {
			pc.Count++
		} else {
			packCounts[g.PackID] = &entity.PackCount{ID: g.PackID, Title: g.PackTitle, Count: 1}
		}
	}

	top := make([]entity.PackCount, 0, len(packCounts))
	for _, pc := range packCounts {
		top = append(top, *pc)
	}
	sort.Slice(top, func(i, j int) bool {
		if top[i].Count != top[j].Count {
			return top[i].Count > top[j].Count
		}
		return top[i].Title < top[j].Title
	})
	if len(top) > topPackLimit {
		top = top[:topPackLimit]
	}

	return &entity.Result{
		Window:   entity.Window{Days: days, From: from, To: to},
		Buckets:  buckets,
		TopPacks: top,
		Summary:  summary,
	}, nil
}

// side describes which row half belongs to the requesting user.
type side struct {
	bingo   int
	numbers []int
}

func pickSide(g entity.GameRow, userID string) side {
	if g.User1ID == userID {
		return side{bingo: g.User1Bingo, numbers: g.User1Numbers}
	}
	return side{bingo: g.User2Bingo, numbers: g.User2Numbers}
}

func countNonZero(nums []int) int {
	n := 0
	for _, v := range nums {
		if v != 0 {
			n++
		}
	}
	return n
}

func clampDays(days int) int {
	if days <= 0 {
		return defaultDays
	}
	if days <= bucketCount7d {
		return bucketCount7d
	}
	if days >= maxDays {
		return maxDays
	}
	return maxDays
}

func bucketCount(days int) int {
	if days <= bucketCount7d {
		return bucketCount7d
	}
	return bucketCount30d
}

// bucketIndex maps a game finish time to a bucket index 0..bucketN-1, where
// bucketN-1 is the most recent bucket (containing `to`).
func bucketIndex(finished time.Time, to time.Time, span time.Duration, bucketN int) int {
	delta := to.Sub(finished)
	if delta < 0 {
		return bucketN - 1
	}
	idx := bucketN - 1 - int(delta/span)
	if idx < 0 {
		return -1
	}
	if idx >= bucketN {
		return bucketN - 1
	}
	return idx
}

// bucketLabel renders the human-friendly tick label for bucket i (0 = oldest).
func bucketLabel(to time.Time, days, bucketN, i int) string {
	span := time.Duration(days) / time.Duration(bucketN) * 24 * time.Hour
	// bucket i covers [to - (bucketN-i)*span, to - (bucketN-1-i)*span]
	start := to.Add(-time.Duration(bucketN-i) * span)
	if bucketN == bucketCount7d {
		return start.Format("Mon")
	}
	return start.Format("Jan 02")
}
