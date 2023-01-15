package userClient

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (u userClient) RequestFriend(login, person string) error {
	_, err := u.conn.RequestFriend(context.Background(), &api.FriendRequest{
		Login:  login,
		Person: person,
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

func (u userClient) AcceptFriend(login, person string) error {
	_, err := u.conn.AcceptFriend(context.Background(), &api.FriendRequest{
		Login:  login,
		Person: person,
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

func (u userClient) DeleteFriend(login, person string) error {
	_, err := u.conn.DeleteFriend(context.Background(), &api.FriendRequest{
		Login:  login,
		Person: person,
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}
