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
	userID, err := h.service.RegisterUser(ctx, req.Username, req.Password, req.Email, req.City)

	switch {
	case err == errs.ErrCredentialsInUse:
		return &api.RegisterUserResponse{}, status.Error(codes.AlreadyExists, "CIU")
	case err != nil:
		return &api.RegisterUserResponse{}, err
	}

	return &api.RegisterUserResponse{UserID: userID, Username: req.Username}, nil
}
