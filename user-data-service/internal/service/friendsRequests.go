package service

import (
	"context"
)

// AcceptFriend writes data when user accepts friendship
func (s service) AcceptFriend(ctx context.Context, userID, friendID string) error {
	return s.repository.AcceptFriend(ctx, userID, friendID)
}

// DeleteFriend writes data when user cancels/declines friendship
func (s service) DeleteFriend(ctx context.Context, userID, friendID string) error {
	return s.repository.DeleteFriend(ctx, userID, friendID)
}

// RequestFriend writes data when user requests friendship
func (s service) RequestFriend(ctx context.Context, userID, friendID string) error {
	return s.repository.RequestFriend(ctx, userID, friendID)
}
