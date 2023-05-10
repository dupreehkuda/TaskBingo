package repository

import (
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func mapToAssignNewPackRequest(userID, packID, packName string) *api.AssignNewPackRequest {
	return &api.AssignNewPackRequest{
		UserID:   userID,
		PackID:   packID,
		PackName: packName,
	}
}

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

func mapFromRatedPacks(packs *api.RatedPacksResponse) *models.Packs {
	var res []string
	for _, val := range packs.Packs {
		res = append(res, val)
	}

	return res
}

func mapFromUserDataResponse(data *api.GetUserDataResponse) *models.UserAccountInfoResponse {
	res := models.UserAccountInfoResponse{
		UserID:     data.UserID,
		Username:   data.Username,
		City:       data.City,
		Wins:       int(data.Wins),
		Lose:       int(data.Loses),
		Bingo:      int(data.Bingo),
		Friends:    []models.FriendsInfo{},
		LikedPacks: data.LikedPacks,
		RatedPacks: data.RatedPacks,
	}

	for _, val := range data.Friends {
		res.Friends = append(res.Friends, models.FriendsInfo{
			UserID:   val.UserID,
			Username: val.Username,
			Status:   int(val.Status),
			Wins:     int(val.Wins),
			Loses:    int(val.Loses),
		})
	}

	return &res
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
