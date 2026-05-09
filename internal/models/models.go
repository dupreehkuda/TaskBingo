package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Game status enumeration. Carry-over from the legacy game-service.
const (
	_ = iota
	GameCreated
	GameWaiting
	GameStart
	GameInProcess
	GameOneFinished
	GameEnd
)

// Friend status enumeration. Carry-over from the legacy user-data-service.
const (
	_ FriendStatus = iota
	FriendRequested
	FriendResponse
	FriendFriend
)

// FriendStatus is the typed enum for the friends.status column. Using a typed
// int avoids re-using the same iota namespace as game statuses.
type FriendStatus int

// Persisted game-status values used in the games table.
const (
	_ DBGameStatus = iota
	GameDBRequested
	GameDBStarted
	GameDBEnded
)

// DBGameStatus is the typed enum for the games.status column.
type DBGameStatus int

// WS heartbeat constants. Preserved so existing clients keep working.
const (
	PongWait     = 5 * time.Second
	PingInterval = (PongWait * 9) / 10
)

// UserIDKey is the typed context-key used by JWT middleware.
type UserIDKey string

const (
	CtxUserIDKey   UserIDKey = "userID"
	CtxUsernameKey UserIDKey = "username"
)

type (
	Users []User
	Packs []TaskPack

	Response struct {
		UserID   string `json:"userID"`
		Username string `json:"username"`
		Points   int    `json:"points"`
		Email    string `json:"email"`
	}

	RegisterCredentials struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		City     string `json:"city"`
		Password string `json:"password"`
	}

	LoginCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Pack struct {
		Title string   `json:"title"`
		Tasks []string `json:"tasks"`
	}

	TaskPack struct {
		ID        string `json:"id"`
		Pack      Pack   `json:"pack"`
		IsPrivate bool   `json:"isPrivate"`
		Creator   string `json:"creator,omitempty"`
	}

	TaskPackRequest struct {
		TaskID string `json:"id"`
	}

	TaskPacksRequest struct {
		PackIDs []string `json:"ids"`
	}

	PackAction struct {
		Pack string `json:"pack"`
	}

	FriendsInfo struct {
		UserID       string    `json:"userID"`
		Username     string    `json:"username"`
		Status       int       `json:"status"`
		Wins         int       `json:"wins"`
		Loses        int       `json:"loses"`
		FriendsSince time.Time `json:"friendsSince"`
	}

	UserAccountInfo struct {
		UserID     string        `json:"userID"`
		Username   string        `json:"username"`
		City       string        `json:"city"`
		Wins       int           `json:"wins"`
		Lose       int           `json:"lose"`
		Bingo      int           `json:"bingo"`
		SoloBingo  int           `json:"soloBingo"`
		Friends    []FriendsInfo `json:"friends"`
		LikedPacks []TaskPack    `json:"likedPacks"`
		RatedPacks []string      `json:"ratedPacks"`
		Games      []GameShort   `json:"games"`
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
		Kind         string  `json:"kind"`
	}

	GameShort struct {
		GameID     string `json:"gameId"`
		User1Id    string `json:"user1Id"`
		User2Id    string `json:"user2Id"`
		PackId     string `json:"packId"`
		Status     int32  `json:"status"`
		User1Bingo int32  `json:"user1Bingo"`
		User2Bingo int32  `json:"user2Bingo"`
		Winner     string `json:"winner"`
		Kind       string `json:"kind"`
	}

	SoloStartRequest struct {
		PackID string `json:"packID"`
	}

	SoloStartResponse struct {
		GameID  string  `json:"gameID"`
		Numbers []int32 `json:"numbers"`
	}

	SoloProgressRequest struct {
		GameID      string  `json:"gameID"`
		UserNumbers []int32 `json:"userNumbers"`
	}

	SoloFinishRequest struct {
		GameID      string  `json:"gameID"`
		UserNumbers []int32 `json:"userNumbers"`
	}

	SoloFinishResponse struct {
		Bingo int32 `json:"bingo"`
	}

	Comment struct {
		ID        string    `json:"id"`
		GameID    string    `json:"gameID"`
		UserID    string    `json:"userID"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	AddCommentRequest struct {
		Body string `json:"body"`
	}

	EditCommentRequest struct {
		Body string `json:"body"`
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
		Status  int     `json:"status"`
		UserID  string  `json:"userID"`
		Bingo   int32   `json:"bingo"`
		Numbers []int32 `json:"userNumbers"`
	}

	Player struct {
		Id       string `json:"id"`
		Finished bool   `json:"finished"`
		Conn     *websocket.Conn
	}

	Room struct {
		Id      string `json:"id"`
		Status  int    `json:"status"`
		Game    *Game
		Player1 *Player `json:"player1"`
		Player2 *Player `json:"player2"`
	}

	GameHub struct {
		Mu    sync.Mutex
		Rooms map[string]*Room
	}
)

// PongHandler resets the read deadline. Mirror of the legacy method.
func (p *Player) PongHandler(_ string) error {
	return p.Conn.SetReadDeadline(time.Now().Add(PongWait))
}
