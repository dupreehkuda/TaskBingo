package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) GetAllPeople(ctx context.Context, _ *api.Empty) (*api.People, error) {
	users, err := h.processor.GetAllUsers()
	if err != nil {
		h.logger.Error("Unable to call processors", zap.Error(err))
		return nil, err
	}

	var resp []*api.PersonInfo

	for _, person := range *users {
		resp = append(resp, &api.PersonInfo{
			Login: person.Login,
			City:  person.City,
		})
	}

	return &api.People{Person: resp}, nil
}
