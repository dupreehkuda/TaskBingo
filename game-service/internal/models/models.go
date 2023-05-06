package models

import (
	"time"
)

type (
	// UserIDKey is type for context keys
	UserIDKey string

	// Response provides user data response
	Response struct {
		UserID   string `json:"userID"`
		Username string `json:"username"`
		Points   int    `json:"points"`
		Email    string `json:"email"`
	}

	// RegisterCredentials provides users register credentials request
	RegisterCredentials struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		City     string `json:"city"`
		Password string `json:"password"`
	}

	// LoginCredentials provides users login credentials request
	LoginCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Pack provides complete task pack info
	Pack struct {
		Title string   `json:"title"`
		Tasks []string `json:"tasks"`
	}

	// TaskPack provides task pack with id
	TaskPack struct {
		ID   string `json:"id"`
		Pack Pack   `json:"pack"`
	}

	// TaskPackRequest provides task pack request
	TaskPackRequest struct {
		TaskID string `json:"id"`
	}

	FriendsInfo struct {
		UserID       string    `json:"userID"`
		Username     string    `json:"username"`
		Status       int       `json:"status"`
		Wins         int       `json:"wins"`
		Loses        int       `json:"loses"`
		FriendsSince time.Time `json:"friendsSince"`
	}

	// UserAccountInfo provides basic account info
	UserAccountInfo struct {
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

	UserAccountInfoResponse struct {
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

	User struct {
		UserID   string `json:"userID"`
		Username string `json:"username"`
		City     string `json:"city"`
		Bingo    int    `json:"bingo"`
	}

	FriendRequest struct {
		Person string `json:"person"`
	}

	// Game provides game instance
	Game struct {
		GameID       string  `json:"gameID"`
		User1Id      string  `json:"user1ID"`
		User2Id      string  `json:"user2ID"`
		PackId       string  `json:"packID"`
		Status       int32   `json:"status"`
		User1Bingo   int32   `json:"user1Bingo"`
		User2Bingo   int32   `json:"user2Bingo"`
		Winner       string  `json:"winner"`
		Numbers      []int32 `json:"numbers"`
		User1Numbers []int32 `json:"user1Numbers"`
		User2Numbers []int32 `json:"user2Numbers"`
	}

	NewGameRequest struct {
		OpponentID string `json:"opponent"`
		Pack       string `json:"pack"`
	}

	StatusGameRequest struct {
		GameID string `json:"gameID"`
	}

	GameAction struct {
		UserID   string  `json:"userID"`
		Finished bool    `json:"finished"`
		Numbers  []int32 `json:"numbers"`
	}

	GameUpdate struct {
		Started bool         `json:"started"`
		Player1 PlayerUpdate `json:"player1"`
		Player2 PlayerUpdate `json:"player2"`
	}

	PlayerUpdate struct {
		UserID   string  `json:"userID"`
		Numbers  []int32 `json:"userNumbers"`
		Finished bool    `json:"finished"`
	}
)
