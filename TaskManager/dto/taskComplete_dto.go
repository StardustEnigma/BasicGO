package dto

import "time"

type TaskComplete struct {
	Title       string `json:"title"`
	TaskId      int    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}