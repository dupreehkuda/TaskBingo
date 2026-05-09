package entity_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/usecases/game_realtime_action/entity"
)

func TestApplyAction_TransitionsCreatedToWaiting(t *testing.T) {
	r := &models.Room{Status: models.GameCreated, Player1: &models.Player{Id: "u1"}}
	require.Nil(t, entity.ApplyAction(r, &models.GameAction{}))
	require.Equal(t, models.GameWaiting, r.Status)
}

func TestApplyAction_StartWhenBothPresent(t *testing.T) {
	r := &models.Room{
		Status:  models.GameWaiting,
		Player1: &models.Player{Id: "u1"},
		Player2: &models.Player{Id: "u2"},
		Game:    &models.Game{User1Id: "u1", User2Id: "u2"},
	}
	update := entity.ApplyAction(r, &models.GameAction{})
	require.NotNil(t, update)
	require.Equal(t, models.GameStart, r.Status)
}

func TestApplyAction_PerUserBingoCount(t *testing.T) {
	r := &models.Room{
		Status:  models.GameInProcess,
		Player1: &models.Player{Id: "u1"},
		Player2: &models.Player{Id: "u2"},
		Game: &models.Game{
			User1Id:      "u1",
			User2Id:      "u2",
			User1Numbers: make([]int32, 16),
			User2Numbers: make([]int32, 16),
		},
	}
	full := []int32{1, 2, 3, 4, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	update := entity.ApplyAction(r, &models.GameAction{UserID: "u1", Numbers: full})
	require.NotNil(t, update)
	require.Equal(t, int32(2), update.Bingo) // two horizontal full rows
}
