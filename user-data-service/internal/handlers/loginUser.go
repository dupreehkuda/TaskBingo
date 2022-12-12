package handlers

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) LoginUser(ctx context.Context, req *api.LoginUserRequest) (*api.LoginUserResponse, error) {
	err := h.processor.LoginUser(req.Login, req.Password)

	switch {
	case err == errs.ErrWrongCredentials:
		return &api.LoginUserResponse{}, status.Error(codes.Unauthenticated, "WC")
	case err != nil:
		return &api.LoginUserResponse{Error: err.Error()}, err
	}

	return &api.LoginUserResponse{}, nil
}
