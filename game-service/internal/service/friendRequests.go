package service

import (
	"context"
)

// RequestFriend calls user repository to request friendship
func (s service) RequestFriend(ctx context.Context, userID, friendID string) error {
	return s.repository.RequestFriend(ctx, userID, friendID)
}

// AcceptFriend calls user repository to accept friendship
func (s service) AcceptFriend(ctx context.Context, userID, friendID string) error {
	return s.repository.AcceptFriend(ctx, userID, friendID)
}

// DeleteFriend calls user repository to delete/cancel friendship
func (s service) DeleteFriend(ctx context.Context, userID, friendID string) error {
	return s.repository.DeleteFriend(ctx, userID, friendID)
}
