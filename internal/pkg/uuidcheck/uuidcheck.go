package uuidcheck

import (
	"github.com/google/uuid"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

// Check validates each id is a non-empty parseable UUID.
func Check(ids ...string) error {
	for _, id := range ids {
		if id == "" {
			return errs.ErrEmptyRequest
		}
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
	}
	return nil
}
