package models

type Response struct {
	UserID string `json:"userID"`
	Points int    `json:"points"`
	Email  string `json:"email"`
}

type Request struct {
	UserID string `json:"userID"`
}
