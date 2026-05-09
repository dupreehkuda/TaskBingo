//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package solo_game_start

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/solo_game_start/entity"
)

type storage interface {
	Create(ctx context.Context, g *models.Game) error
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

// Run constructs a solo game via entity.NewSoloGame and persists it. Returns
// the game ID and the shuffled number set the client renders on the board.
func (u *Usecase) Run(ctx context.Context, userID, packID string) (*models.SoloStartResponse, error) {
	g, err := entity.NewSoloGame(userID, packID)
	if err != nil {
		return nil, err
	}
	if err := u.tx.RunInTx(ctx, func(ctx context.Context) error {
		return u.storage.Create(ctx, g)
	}); err != nil {
		u.logger.Error("solo create failed", zap.Error(err))
		return nil, err
	}
	return &models.SoloStartResponse{GameID: g.GameID, Numbers: g.Numbers}, nil
}
