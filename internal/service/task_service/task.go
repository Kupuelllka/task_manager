package task

import "time"

// Task представляет собой структуру задачи
type Task struct {
	ID         string        `json:"id"`
	Status     string        `json:"status"` // "pending", "processing", "completed", "failed"
	CreatedAt  time.Time     `json:"created_at"`
	StartedAt  time.Time     `json:"started_at,omitempty"`
	FinishedAt time.Time     `json:"finished_at,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	Result     string        `json:"result,omitempty"`
}
