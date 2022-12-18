package models

// GetUserDataResponse provides response for user data request
type GetUserDataResponse struct {
	UserID string `json:"userID"`
	Points int    `json:"points"`
	Email  string `json:"email"`
}

// LoginUserResponse provides response for user login request
type LoginUserResponse struct {
	PasswordHash string `db:"passwordhash"`
	PasswordSalt string `db:"passwordsalt"`
}
