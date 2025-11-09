package types

import "time"

type CreateTask struct {
	ProjectID   string    `json:"-" validate:"required,uuid"`
	Title       string    `json:"title" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=1024"`
	Status      string    `json:"-" validate:"required,oneof=todo in_progress in_review done"`
	Priority    string    `json:"priority" validate:"required,oneof=low medium high urgent"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	AssignedTo  *string   `json:"assigned_to" validate:"omitempty,uuid"`
	CreatedBy   string    `json:"-" validate:"required,uuid"`
	DependsOn   []string  `json:"depends_on" validate:"omitempty,dive,required,uuid"`
}

type ListTasks struct {
	ProjectID string `json:"-" validate:"required,uuid"`
	Filter    string `json:"filter" validate:"required,oneof=created_at due_date priority"`
	Direction string `json:"direction" validate:"required,oneof=asc desc"`
	Page      int    `json:"page" validate:"required,min=1"`
	Limit     int    `json:"limit" validate:"required,min=1,max=100"`
}

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
