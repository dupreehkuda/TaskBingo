//go:generate easyjson -no_std_marshalers models.go
package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	_ = iota
	GameCreated
	GameWaiting
	GameStart
	GameInProcess
	GameOneFinished
	GameEnd
)

const (
	// PongWait is how long we will await a pong response from client
	PongWait = 5 * time.Second

	// PingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
	PingInterval = (PongWait * 9) / 10
)

// UserIDKey is type for context keys
type UserIDKey string

//easyjson:json
type (
	Users []User
	Packs []TaskPack

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

	// TaskPacksRequest provides task pack request
	TaskPacksRequest struct {
		PackIDs []string `json:"ids"`
	}

	// FriendsInfo provides data about user's friend
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
		Games      []GameShort   `json:"games"`
	}

	// User provides brief user info
	User struct {
		UserID   string `json:"userID"`
		Username string `json:"username"`
		City     string `json:"city"`
		Bingo    int    `json:"bingo"`
	}

	// FriendRequest contains userID needed for friendship requests
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

	// GameShort provides brief game info
	GameShort struct {
		GameID     string `json:"gameId"`
		User1Id    string `json:"user1Id"`
		User2Id    string `json:"user2Id"`
		PackId     string `json:"packId"`
		Status     int32  `json:"status"`
		User1Bingo int32  `json:"user1Bingo"`
		User2Bingo int32  `json:"user2Bingo"`
		Winner     string `json:"winner"`
	}

	// NewGameRequest provides info for new game
	NewGameRequest struct {
		OpponentID string `json:"opponent"`
		Pack       string `json:"pack"`
	}

	// StatusGameRequest provides gameID for accepting/deleting a game
	StatusGameRequest struct {
		GameID string `json:"gameID"`
	}

	// GameAction provides info of incoming action
	GameAction struct {
		UserID   string  `json:"userID"`
		Finished bool    `json:"finished"`
		Numbers  []int32 `json:"numbers"`
	}

	// GameUpdate provides response for incoming action
	GameUpdate struct {
		Status  int     `json:"status"`
		UserID  string  `json:"userID"`
		Bingo   int32   `json:"bingo"`
		Numbers []int32 `json:"userNumbers"`
	}

	// Player provides info about connected user
	Player struct {
		Id       string `json:"id"`
		Finished bool   `json:"finished"`
		Conn     *websocket.Conn
	}

	// Room provides info about Game and both Player
	Room struct {
		Id      string `json:"id"`
		Status  int    `json:"status"`
		Game    *Game
		Player1 *Player `json:"player1"`
		Player2 *Player `json:"player2"`
	}

	// GameHub stores all active Rooms
	GameHub struct {
		Mu    sync.Mutex
		Rooms map[string]*Room
	}
)

// very stupid implementation, should refactor with client struct
func (p *Player) PongHandler(appdata string) error {
	return p.Conn.SetReadDeadline(time.Now().Add(PongWait))
}
