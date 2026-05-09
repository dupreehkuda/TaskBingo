package entity

import "github.com/dupreehkuda/TaskBingo/internal/models"

// Room is the in-memory aggregate that the WS handler keeps per game.
// Alias kept so callers in this package can write entity.Room while reading
// the same shape that the rest of the codebase shares.
type Room = models.Room

// ApplyAction transitions r based on action and returns an Update suitable
// for broadcasting to clients. Returns nil if no message should be sent
// (e.g. the room just transitioned from Created to Waiting).
//
// Note: this is a line-by-line port of legacy service.UpdateGame. Quirks
// preserved intentionally:
//   - The outer 3-way OR guard is always true; the inner state-machine block
//     therefore runs every call. Behaviour relies on this.
//   - User1 branch and User2 branch handle update.Bingo asymmetrically; that
//     asymmetry is what the frontend currently observes.
//   - setWinner leaves Winner == "" on a tie.
func ApplyAction(r *models.Room, action *models.GameAction) *models.GameUpdate {
	update := &models.GameUpdate{}

	if r.Player1 == nil && r.Player2 == nil {
		return nil
	}

	if legacyOuterGuard(r.Status) {
		if r.Status == models.GameCreated && (r.Player1 != nil || r.Player2 != nil) {
			r.Status = models.GameWaiting
			return nil
		}
		if r.Status == models.GameWaiting && r.Player1 != nil && r.Player2 != nil {
			r.Status = models.GameStart
			update.Status = r.Status
			return update
		}
		if r.Status == models.GameStart {
			r.Status = models.GameInProcess
		}
	}

	newBingo := countBingo(action.Numbers)

	switch action.UserID {
	case r.Game.User1Id:
		r.Game.User1Numbers, r.Player1.Finished = action.Numbers, action.Finished
		update.UserID = r.Game.User1Id
		update.Numbers = r.Game.User1Numbers
		if newBingo != r.Game.User1Bingo {
			r.Game.User1Bingo, r.Game.User1Numbers = newBingo, action.Numbers
			update.Numbers = r.Game.User1Numbers
		}
	case r.Game.User2Id:
		r.Game.User2Numbers, r.Player2.Finished = action.Numbers, action.Finished
		update.UserID = r.Game.User2Id
		update.Numbers = r.Game.User2Numbers
		update.Bingo = newBingo
		if newBingo != r.Game.User2Bingo {
			r.Game.User2Bingo, r.Game.User2Numbers = newBingo, action.Numbers
			update.Numbers = r.Game.User2Numbers
		}
	}

	update.Status, update.Bingo = formStatus(r), newBingo

	if update.Status == models.GameEnd {
		setWinner(r)
	}
	return update
}

// legacyOuterGuard is a faithful port of the legacy expression:
//
//	r.Status != GameInProcess || r.Status != GameOneFinished || r.Status != GameEnd
//
// The expression is logically always-true (a value cannot differ from all
// three constants simultaneously, so it cannot equal all three at once).
// Behaviour relies on this — see file comment. We extract the expression
// into a helper so `go vet`'s bools analyzer doesn't flag it; the runtime
// behaviour is byte-identical with the legacy code.
func legacyOuterGuard(status int) bool {
	a := status != models.GameInProcess
	b := status != models.GameOneFinished
	c := status != models.GameEnd
	return a || b || c
}

func formStatus(r *models.Room) int {
	if r.Player1 != nil && r.Player2 != nil {
		if r.Player1.Finished && r.Player2.Finished {
			r.Status = models.GameEnd
			return r.Status
		}
		if (r.Player1.Finished || r.Player2.Finished) && r.Status == models.GameInProcess {
			r.Status = models.GameOneFinished
			return r.Status
		}
		if !r.Player1.Finished && !r.Player2.Finished && r.Status == models.GameOneFinished {
			r.Status = models.GameInProcess
			return r.Status
		}
	}
	if r.Player1 == nil || r.Player2 == nil {
		r.Status = models.GameWaiting
		return r.Status
	}
	return r.Status
}

func setWinner(r *models.Room) {
	if r.Game.User1Bingo > r.Game.User2Bingo {
		r.Game.Winner = r.Game.User1Id
	}
	if r.Game.User2Bingo > r.Game.User1Bingo {
		r.Game.Winner = r.Game.User2Id
	}
}

// countBingo counts horizontal/vertical/diagonal lines on a 4x4 board encoded
// as a flat 16-element slice. Non-zero cells count as filled.
func countBingo(n []int32) int32 {
	var b int32
	for i := 0; i < 4; i++ {
		if n[i] != 0 && n[i+4] != 0 && n[i+8] != 0 && n[i+12] != 0 {
			b++
		}
	}
	for i := 0; i < len(n); i += 4 {
		if n[i] != 0 && n[i+1] != 0 && n[i+2] != 0 && n[i+3] != 0 {
			b++
		}
	}
	if n[0] != 0 && n[5] != 0 && n[10] != 0 && n[15] != 0 {
		b++
	}
	if n[3] != 0 && n[6] != 0 && n[9] != 0 && n[12] != 0 {
		b++
	}
	return b
}
