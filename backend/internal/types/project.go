package types

import "time"

type CreateProject struct {
	UserID      string `json:"-" validate:"-"`
	WorkspaceID string `json:"-" validate:"-"`
	Name        string `json:"name" validate:"required,max=32"`
	Description string `json:"description" validate:"required,max=1024"`
	Color       string `json:"color" validate:"required,hexcolor"`
	Status      string `json:"status" validate:"required,oneof=active archived completed"`
}

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	WorkspaceID string    `json:"workspace_id"`
}
