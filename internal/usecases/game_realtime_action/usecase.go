//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package game_realtime_action

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_realtime_action/entity"
)

type storage interface {
	Get(ctx context.Context, gameID string) (*models.Game, error)
	Achieve(ctx context.Context, g *models.Game) error
}

type txManager interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Usecase struct {
	storage storage
	tx      txManager
	logger  *zap.Logger
}

func New(s storage, tx txManager, logger *zap.Logger) *Usecase {
	return &Usecase{storage: s, tx: tx, logger: logger}
}

// LoadRoom builds an empty Room around the persisted game. The hub uses this
// only on the first connection to a given gameID.
func (u *Usecase) LoadRoom(ctx context.Context, gameID string) (*models.Room, error) {
	g, err := u.storage.Get(ctx, gameID)
	if err != nil {
		return nil, err
	}
	return &models.Room{Id: gameID, Game: g, Status: models.GameCreated}, nil
}

// Run applies action to room and, on GameEnd, persists final state.
// Returns the update to be broadcast (or nil if no broadcast).
func (u *Usecase) Run(ctx context.Context, r *models.Room, action *models.GameAction) (*models.GameUpdate, error) {
	update := entity.ApplyAction(r, action)
	if update == nil {
		return nil, nil
	}
	if update.Status == models.GameEnd {
		if err := u.tx.RunInTx(ctx, func(ctx context.Context) error {
			return u.storage.Achieve(ctx, r.Game)
		}); err != nil {
			u.logger.Error("achieve failed", zap.Error(err))
			return update, err
		}
	}
	return update, nil
}
