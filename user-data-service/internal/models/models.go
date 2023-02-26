package models

import (
	"time"

	"github.com/google/uuid"
)

type FriendsInfo struct {
	Login        string    `json:"login"`
	Status       int       `json:"status"`
	Wins         int       `json:"wins"`
	Loses        int       `json:"loses"`
	FriendsSince time.Time `json:"friendsSince"`
}

// GetUserDataResponse provides response for user data request
type GetUserDataResponse struct {
	Login      string        `json:"login"`
	City       string        `json:"city"`
	Wins       int           `json:"wins"`
	Lose       int           `json:"lose"`
	Bingo      int           `json:"bingo"`
	Friends    []FriendsInfo `json:"friends"`
	LikedPacks []string      `json:"likedPacks"`
	RatedPacks []string      `json:"ratedPacks"`
}

// LoginUserResponse provides response for user login request
type LoginUserResponse struct {
	PasswordHash string `db:"passwordhash"`
	PasswordSalt string `db:"passwordsalt"`
}

type AllUsers struct {
	Login string `json:"login"`
	City  string `json:"city"`
	Bingo int    `json:"bingo"`
}

// Game provides game instance
type Game struct {
	GameID       uuid.UUID `json:"gameId"`
	User1Id      string    `json:"user1Id"`
	User2Id      string    `json:"user2Id"`
	PackId       string    `json:"packId"`
	Status       int32     `json:"status"`
	User1Bingo   int32     `json:"user1Bingo"`
	User2Bingo   int32     `json:"user2Bingo"`
	Winner       string    `json:"winner"`
	Numbers      []int32   `json:"numbers"`
	User1Numbers []int32   `json:"user1Numbers"`
	User2Numbers []int32   `json:"user2Numbers"`
}
