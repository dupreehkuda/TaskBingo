package models

type GetUserDataResponse struct {
	UserID string `json:"userID"`
	Points int    `json:"points"`
	Email  string `json:"email"`
}

type LoginUserResponse struct {
	PasswordHash string `db:"passwordhash"`
	PasswordSalt string `db:"passwordsalt"`
}
