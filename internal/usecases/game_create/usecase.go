//go:generate go tool mockgen -source=usecase.go -destination=mocks/usecase.go -package=mocks

package game_create

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_create/entity"
)

type storage interface {
	Create(ctx context.Context, game *models.Game) error
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

// Run constructs the game via entity.NewGame and persists it.
// Returns the GameShort projection that the legacy handler used to send back.
func (u *Usecase) Run(ctx context.Context, userID, opponentID, packID string) (*models.GameShort, error) {
	g, err := entity.NewGame(userID, opponentID, packID)
	if err != nil {
		return nil, err
	}
	if err := u.tx.RunInTx(ctx, func(ctx context.Context) error {
		return u.storage.Create(ctx, g)
	}); err != nil {
		return nil, err
	}
	return &models.GameShort{
		GameID:  g.GameID,
		User1Id: g.User1Id,
		User2Id: g.User2Id,
		PackId:  g.PackId,
		Status:  g.Status,
	}, nil
}
