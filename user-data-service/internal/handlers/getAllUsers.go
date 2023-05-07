package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) GetAllPeople(ctx context.Context, _ *api.Empty) (*api.People, error) {
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	var resp []*api.PersonInfo

	for _, person := range *users {
		resp = append(resp, &api.PersonInfo{
			UserID:   person.UserID,
			Username: person.Username,
			City:     person.City,
			Bingo:    int32(person.Bingo),
		})
	}

	return &api.People{Person: resp}, nil
}
