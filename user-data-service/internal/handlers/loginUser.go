package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// LoginUser handles the operation of user's logging in
func (h *Handlers) LoginUser(ctx context.Context, req *api.LoginUserRequest) (*api.LoginUserResponse, error) {
	userID, err := h.processor.LoginUser(req.Username, req.Password)

	switch {
	case err == errs.ErrWrongCredentials:
		return &api.LoginUserResponse{}, status.Error(codes.Unauthenticated, "Wrong Credentials")
	case err != nil:
		return &api.LoginUserResponse{}, err
	}

	return &api.LoginUserResponse{
		UserID:   &api.UUID{Id: userID},
		Username: req.Username,
	}, nil
}
