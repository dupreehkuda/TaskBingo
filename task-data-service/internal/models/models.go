package models

// TaskPack provides complete task pack info
type TaskPack struct {
	TaskID string   `json:"id"`
	Tasks  []string `json:"tasks"`
}
