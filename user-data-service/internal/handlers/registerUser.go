package handlers

import (
	"context"
	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	err := h.processor.RegisterUser(req.Login, req.Password, req.Email)
	if err == errs.ErrCredentialsInUse {
		return &api.RegisterUserResponse{Error: "CIU"}, nil
	}

	switch {
	case err == errs.ErrCredentialsInUse:
		return &api.RegisterUserResponse{Error: "CIU"}, nil
	case err != nil:
		return &api.RegisterUserResponse{Error: err.Error()}, err
	}

	return &api.RegisterUserResponse{}, nil
}
