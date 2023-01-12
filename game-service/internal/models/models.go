package models

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
	Login string `json:"login"`
	City  string `json:"city"`
	Bingo string `json:"bingo"`
}

// UserAccountInfo provides basic account info
type UserAccountInfo struct {
	Login   string        `json:"login"`
	City    string        `json:"city"`
	Wins    int           `json:"wins"`
	Lose    int           `json:"lose"`
	Bingo   int           `json:"bingo"`
	Friends []FriendsInfo `json:"friends"`
	Packs   []TaskPack    `json:"packs"`
}

type UserAccountInfoResponse struct {
	Login   string        `json:"login"`
	City    string        `json:"city"`
	Wins    int           `json:"wins"`
	Lose    int           `json:"lose"`
	Bingo   int           `json:"bingo"`
	Friends []FriendsInfo `json:"friends"`
	Packs   []string      `json:"packs"`
}

type User struct {
	Login string `json:"login"`
	City  string `json:"city"`
	Bingo int    `json:"bingo"`
}
