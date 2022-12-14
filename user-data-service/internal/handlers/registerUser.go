package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// RegisterUser handles the operation of user's registration
func (h *Handlers) RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	err := h.processor.RegisterUser(req.Login, req.Password, req.Email, req.City)

	switch {
	case err == errs.ErrCredentialsInUse:
		return &api.RegisterUserResponse{}, status.Error(codes.AlreadyExists, "CIU")
	case err != nil:
		return &api.RegisterUserResponse{Error: err.Error()}, err
	}

	return &api.RegisterUserResponse{}, nil
}
