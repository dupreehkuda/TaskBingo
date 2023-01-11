package models

type FriendsInfo struct {
	Login string `json:"login"`
	City  string `json:"city"`
	Bingo string `json:"bingo"`
}

// GetUserDataResponse provides response for user data request
type GetUserDataResponse struct {
	Login   string        `json:"login"`
	City    string        `json:"city"`
	Wins    int           `json:"wins"`
	Lose    int           `json:"lose"`
	Bingo   int           `json:"bingo"`
	Friends []FriendsInfo `json:"friends"`
	Packs   []string      `json:"packs"`
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
