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
