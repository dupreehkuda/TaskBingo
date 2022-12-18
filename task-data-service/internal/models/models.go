package models

type TaskPack struct {
	TaskID string   `json:"id"`
	Tasks  []string `json:"tasks"`
}
