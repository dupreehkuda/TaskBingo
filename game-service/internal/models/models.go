package models

import (
	"time"

	"github.com/google/uuid"
)

// LoginKey is type for context keys
type LoginKey string

// Response provides user data response
type Response struct {
	UserID string `json:"userID"`
	Points int    `json:"points"`
	Email  string `json:"email"`
}

// Request provides user data request
type Request struct {
	UserID string `json:"userID"`
}

// RegisterCredentials provides users register credentials request
type RegisterCredentials struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	City     string `json:"city"`
	Password string `json:"password"`
}

// LoginCredentials provides users login credentials request
type LoginCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthResponse provides users authorization response
type AuthResponse struct {
	Login string `json:"login"`
	Token string `json:"token"`
}

// TaskPack provides complete task pack ingo
type TaskPack struct {
	TaskID string   `json:"id"`
	Tasks  []string `json:"tasks"`
}

// TaskPackRequest provides task pack request
type TaskPackRequest struct {
	TaskID string `json:"id"`
}

type FriendsInfo struct {
	Login        string    `json:"login"`
	Status       int       `json:"status"`
	Wins         int       `json:"wins"`
	Loses        int       `json:"loses"`
	FriendsSince time.Time `json:"friendsSince"`
}

// UserAccountInfo provides basic account info
type UserAccountInfo struct {
	Login      string        `json:"login"`
	City       string        `json:"city"`
	Wins       int           `json:"wins"`
	Lose       int           `json:"lose"`
	Bingo      int           `json:"bingo"`
	Friends    []FriendsInfo `json:"friends"`
	LikedPacks []TaskPack    `json:"likedPacks"`
	RatedPacks []string      `json:"ratedPacks"`
}

type UserAccountInfoResponse struct {
	Login      string        `json:"login"`
	City       string        `json:"city"`
	Wins       int           `json:"wins"`
	Lose       int           `json:"lose"`
	Bingo      int           `json:"bingo"`
	Friends    []FriendsInfo `json:"friends"`
	LikedPacks []string      `json:"likedPacks"`
	RatedPacks []string      `json:"ratedPacks"`
}

type User struct {
	Login string `json:"login"`
	City  string `json:"city"`
	Bingo int    `json:"bingo"`
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
	User string `json:"user"`
	Pack string `json:"pack"`
}
