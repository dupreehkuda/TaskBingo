package repository

import (
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func mapToGameRequest(game *models.Game) *api.GameRequest {
	return &api.GameRequest{
		GameID:       game.GameID,
		User1Id:      game.User1Id,
		User2Id:      game.User2Id,
		Pack:         game.PackId,
		Status:       game.Status,
		User1Bingo:   game.User1Bingo,
		User2Bingo:   game.User2Bingo,
		Winner:       game.Winner,
		Numbers:      game.Numbers,
		User1Numbers: game.User1Numbers,
		User2Numbers: game.User2Numbers,
	}
}

func mapFromGameRequest(game *api.GameRequest) *models.Game {
	return &models.Game{
		GameID:       game.GameID,
		User1Id:      game.User1Id,
		User2Id:      game.User2Id,
		PackId:       game.Pack,
		Status:       game.Status,
		User1Bingo:   game.User1Bingo,
		User2Bingo:   game.User2Bingo,
		Winner:       game.Winner,
		Numbers:      game.Numbers,
		User1Numbers: game.User1Numbers,
		User2Numbers: game.User2Numbers,
	}
}

func mapToStatusGameRequest(userID, gameID string) *api.StatusGameRequest {
	return &api.StatusGameRequest{
		UserID: userID,
		GameID: gameID,
	}
}

func mapToFriendRequest(userID, friendID string) *api.FriendRequest {
	return &api.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	}
}

func mapFromPeople(people *api.People) *models.Users {
	var users models.Users

	for _, person := range people.Person {
		users = append(users, models.User{
			UserID:   person.UserID,
			Username: person.Username,
			City:     person.City,
			Bingo:    int(person.Bingo),
		})
	}

	return &users
}

func mapFromRatedPacks(packs *api.GetMultiplePacksResponse) *models.Packs {
	var res models.Packs

	for _, val := range packs.Packs {
		res = append(res, models.TaskPack{
			ID: val.Id,
			Pack: models.Pack{
				Title: val.Pack.Title,
				Tasks: val.Pack.Tasks,
			},
		})
	}

	return &res
}

func mapFromTaskPack(pack *api.TaskPackResponse) *models.TaskPack {
	return &models.TaskPack{
		ID: pack.Id,
		Pack: models.Pack{
			Title: pack.Pack.Title,
			Tasks: pack.Pack.Tasks,
		},
	}
}

func mapFromUserDataResponse(data *api.GetUserDataResponse) *models.UserAccountInfo {
	res := models.UserAccountInfo{
		UserID:     data.UserID,
		Username:   data.Username,
		City:       data.City,
		Wins:       int(data.Wins),
		Lose:       int(data.Loses),
		Bingo:      int(data.Bingo),
		Friends:    []models.FriendsInfo{},
		LikedPacks: []models.TaskPack{},
		RatedPacks: append([]string{}, data.RatedPacks...),
		Games:      []models.GameShort{},
	}

	for _, val := range data.Friends {
		res.Friends = append(res.Friends, mapToFriend(val))
	}

	for _, val := range data.LikedPacks {
		res.LikedPacks = append(res.LikedPacks, mapToPack(val))
	}

	for _, val := range data.Games {
		res.Games = append(res.Games, mapToGame(val))
	}

	return &res
}

func mapToGame(game *api.GameShort) models.GameShort {
	return models.GameShort{
		GameID:     game.GameID,
		User1Id:    game.User1Id,
		User2Id:    game.User2Id,
		PackId:     game.Pack,
		Status:     game.Status,
		User1Bingo: game.User1Bingo,
		User2Bingo: game.User2Bingo,
		Winner:     game.Winner,
	}
}

func mapToFriend(friend *api.FriendInfo) models.FriendsInfo {
	return models.FriendsInfo{
		UserID:   friend.UserID,
		Username: friend.Username,
		Status:   int(friend.Status),
		Wins:     int(friend.Wins),
		Loses:    int(friend.Loses),
	}
}

func mapToPack(pack *api.TaskPackResponse) models.TaskPack {
	return models.TaskPack{
		ID: pack.Id,
		Pack: models.Pack{
			Title: pack.Pack.Title,
			Tasks: pack.Pack.Tasks,
		},
	}
}

func mapToLikeOrRate(userID, pack string) *api.LikeOrRatePackRequest {
	return &api.LikeOrRatePackRequest{
		UserID: userID,
		Pack:   pack,
	}
}

func mapToLogin(username, password string) *api.LoginUserRequest {
	return &api.LoginUserRequest{
		Username: username,
		Password: password,
	}
}

func mapToRegister(credits *models.RegisterCredentials) *api.RegisterUserRequest {
	return &api.RegisterUserRequest{
		Username: credits.Username,
		Email:    credits.Email,
		City:     credits.City,
		Password: credits.Password,
	}
}

func mapToNewTaskPack(userID string, pack *models.TaskPack) *api.NewTaskPackRequest {
	return &api.NewTaskPackRequest{
		UserID: userID,
		PackID: pack.ID,
		Pack: &api.Pack{
			Title: pack.Pack.Title,
			Tasks: pack.Pack.Tasks,
		},
	}
}
