package models

type FriendsInfo struct {
	Login string `json:"login"`
	City  string `json:"city"`
	Score string `json:"score"`
}

// GetUserDataResponse provides response for user data request
type GetUserDataResponse struct {
	Login      string        `json:"login"`
	City       string        `json:"city"`
	Wins       int           `json:"wins"`
	Lose       int           `json:"lose"`
	Scoreboard int           `json:"scoreboard"`
	Friends    []FriendsInfo `json:"friends"`
	Packs      []string      `json:"packs"`
}

// LoginUserResponse provides response for user login request
type LoginUserResponse struct {
	PasswordHash string `db:"passwordhash"`
	PasswordSalt string `db:"passwordsalt"`
}

type AllUsers struct {
	Login string `json:"login"`
	City  string `json:"city"`
}
