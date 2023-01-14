package models

import "time"

type FriendsInfo struct {
	Login        string    `json:"login"`
	City         string    `json:"city"`
	Bingo        int       `json:"bingo"`
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
