package models

// Pack provides complete task pack info
type Pack struct {
	Title string   `json:"title"`
	Tasks []string `json:"tasks"`
}

// TaskPack provides task pack with id
type TaskPack struct {
	ID   string `json:"id"`
	Pack Pack   `json:"pack"`
}
