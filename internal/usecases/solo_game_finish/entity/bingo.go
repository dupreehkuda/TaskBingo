package entity

// CountBingo counts horizontal/vertical/diagonal lines on a 4x4 board encoded
// as a flat 16-element slice. Non-zero cells count as filled.
//
// Note: indices are byte-identical to game_realtime_action/entity.countBingo.
// Duplicated here per the per-usecase storage rule (no shared component).
func CountBingo(n []int32) int32 {
	if len(n) != 16 {
		return 0
	}
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
