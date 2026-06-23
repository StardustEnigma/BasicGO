package models

import(
	"time"
)
type TaskStatus string

const(
	StatusPending TaskStatus="Pending"
	StatusProgress TaskStatus="In_Progress"
	StatusCompleted TaskStatus="Completed"
)
type Task struct{
	TaskId int `json:"task_id"`
	Title string `json:"title"`
	TaskStatus TaskStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	CompletedAt *time.Time  `json:"completed_at"`
}