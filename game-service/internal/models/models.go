package models

type LoginKey string

type Response struct {
	UserID string `json:"userID"`
	Points int    `json:"points"`
	Email  string `json:"email"`
}

type Request struct {
	UserID string `json:"userID"`
}

type RegisterCredentials struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Login string `json:"login"`
	Token string `json:"token"`
}
