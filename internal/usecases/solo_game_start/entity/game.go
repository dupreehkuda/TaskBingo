package entity

import (
	"math/rand"

	"github.com/google/uuid"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

// NewSoloGame returns a freshly constructed solo Game ready to be persisted.
// User1 is the player; User2 stays empty (storage writes NULL into the column).
func NewSoloGame(userID, packID string) (*models.Game, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &models.Game{
		GameID:       id.String(),
		User1Id:      userID,
		PackId:       packID,
		Status:       0,
		Numbers:      shuffled1to16(),
		User1Numbers: zero16(),
		Kind:         "solo",
	}, nil
}

func shuffled1to16() []int32 {
	n := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	rand.Shuffle(len(n), func(i, j int) { n[i], n[j] = n[j], n[i] })
	return n
}

func zero16() []int32 { return make([]int32, 16) }
