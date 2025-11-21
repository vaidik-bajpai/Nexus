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

type ProjectContextKey string

var ProjectCtxKey = ProjectContextKey("project")

type ProjectList struct {
	SearchTerm string `json:"search_term" validate:"omitempty,max=255"`
	Page       int    `json:"page" validate:"required,number,min=1,gte=1"`
	PageSize   int    `json:"page_size" validate:"required,number,gt=0"`
	Filter     int    `json:"filter" validate:"omitempty"`
	Direction  int    `json:"direction" validate:"required,oneof=asc desc"`
}
