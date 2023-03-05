package models

import (
	"time"

	"github.com/google/uuid"
)

// UserIDKey is type for context keys
type UserIDKey string

// Response provides user data response
type Response struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Points   int    `json:"points"`
	Email    string `json:"email"`
}

// RegisterCredentials provides users register credentials request
type RegisterCredentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	City     string `json:"city"`
	Password string `json:"password"`
}

// LoginCredentials provides users login credentials request
type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TaskPack provides complete task pack ingo
type TaskPack struct {
	Name  string   `json:"Name"`
	Tasks []string `json:"tasks"`
}

// TaskPackRequest provides task pack request
type TaskPackRequest struct {
	TaskID string `json:"id"`
}

type FriendsInfo struct {
	UserID       string    `json:"userID"`
	Username     string    `json:"username"`
	Status       int       `json:"status"`
	Wins         int       `json:"wins"`
	Loses        int       `json:"loses"`
	FriendsSince time.Time `json:"friendsSince"`
}

// UserAccountInfo provides basic account info
type UserAccountInfo struct {
	UserID     string        `json:"userID"`
	Username   string        `json:"username"`
	City       string        `json:"city"`
	Wins       int           `json:"wins"`
	Lose       int           `json:"lose"`
	Bingo      int           `json:"bingo"`
	Friends    []FriendsInfo `json:"friends"`
	LikedPacks []TaskPack    `json:"likedPacks"`
	RatedPacks []string      `json:"ratedPacks"`
}

type UserAccountInfoResponse struct {
	UserID     string        `json:"userID"`
	Username   string        `json:"username"`
	City       string        `json:"city"`
	Wins       int           `json:"wins"`
	Lose       int           `json:"lose"`
	Bingo      int           `json:"bingo"`
	Friends    []FriendsInfo `json:"friends"`
	LikedPacks []string      `json:"likedPacks"`
	RatedPacks []string      `json:"ratedPacks"`
}

type User struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	City     string `json:"city"`
	Bingo    int    `json:"bingo"`
}

type FriendRequest struct {
	Person string `json:"person"`
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

type NewGameRequest struct {
	OpponentID string `json:"opponentID"`
	Pack       string `json:"pack"`
}
