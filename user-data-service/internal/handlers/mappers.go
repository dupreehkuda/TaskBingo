package handlers

import (
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

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

func mapToPeople(users *[]models.AllUsers) *api.People {
	var resp []*api.PersonInfo

	for _, person := range *users {
		resp = append(resp, &api.PersonInfo{
			UserID:   person.UserID,
			Username: person.Username,
			City:     person.City,
			Bingo:    int32(person.Bingo),
		})
	}

	return &api.People{Person: resp}
}

func mapToUserDataResponse(data *models.GetUserDataResponse) *api.GetUserDataResponse {
	resp := &api.GetUserDataResponse{
		UserID:     data.UserID,
		Username:   data.Username,
		City:       data.City,
		Wins:       int32(data.Wins),
		Loses:      int32(data.Lose),
		Bingo:      int32(data.Bingo),
		Friends:    []*api.FriendInfo{},
		LikedPacks: []*api.TaskPackResponse{},
		RatedPacks: data.RatedPacks,
		Games:      []*api.GameShort{},
	}

	for _, val := range data.Friends {
		resp.Friends = append(resp.Friends, mapToFriend(&val))
	}

	for _, val := range data.LikedPacks {
		resp.LikedPacks = append(resp.LikedPacks, mapToPack(&val))
	}

	for _, val := range data.GamesShort {
		resp.Games = append(resp.Games, mapToGame(&val))
	}

	return resp
}

func mapToGame(game *models.GameShort) *api.GameShort {
	return &api.GameShort{
		GameID:     game.GameID,
		User1Id:    game.User1Id,
		User2Id:    game.User2Id,
		Pack:       game.PackId,
		Status:     game.Status,
		User1Bingo: game.User1Bingo,
		User2Bingo: game.User2Bingo,
		Winner:     game.Winner,
	}
}

func mapToFriend(friend *models.FriendsInfo) *api.FriendInfo {
	return &api.FriendInfo{
		UserID:   friend.UserID,
		Username: friend.Username,
		Status:   int32(friend.Status),
		Wins:     int32(friend.Wins),
		Loses:    int32(friend.Loses),
	}
}

func mapToLoginUserResponse(userID, username string) *api.LoginUserResponse {
	return &api.LoginUserResponse{
		UserID:   userID,
		Username: username,
	}
}

func mapToRegisterUserResponse(userID, username string) *api.RegisterUserResponse {
	return &api.RegisterUserResponse{
		UserID:   userID,
		Username: username,
	}
}

func mapToMultiplePacks(packs *[]models.TaskPack) []*api.TaskPackResponse {
	var result []*api.TaskPackResponse

	for _, pack := range *packs {
		result = append(result, &api.TaskPackResponse{
			Id: pack.ID,
			Pack: &api.Pack{
				Title: pack.Pack.Title,
				Tasks: pack.Pack.Tasks,
			},
		})
	}

	return result
}

func mapToPack(pack *models.TaskPack) *api.TaskPackResponse {
	return &api.TaskPackResponse{
		Id: pack.ID,
		Pack: &api.Pack{
			Title: pack.Pack.Title,
			Tasks: pack.Pack.Tasks,
		},
	}
}
