package userClient

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (u userClient) RequestFriend(userID, friendID string) error {
	_, err := u.conn.RequestFriend(context.Background(), &api.FriendRequest{
		UserID:   &api.UUID{Id: userID},
		FriendID: &api.UUID{Id: friendID},
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

func (u userClient) AcceptFriend(userID, friendID string) error {
	_, err := u.conn.AcceptFriend(context.Background(), &api.FriendRequest{
		UserID:   &api.UUID{Id: userID},
		FriendID: &api.UUID{Id: friendID},
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

func (u userClient) DeleteFriend(userID, friendID string) error {
	_, err := u.conn.DeleteFriend(context.Background(), &api.FriendRequest{
		UserID:   &api.UUID{Id: userID},
		FriendID: &api.UUID{Id: friendID},
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}
