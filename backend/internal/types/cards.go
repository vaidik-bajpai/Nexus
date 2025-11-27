package types

import "time"

type CreateCard struct {
	ListID  string `json:"-" validate:"required,uuid"`
	UserID  string `json:"-" validate:"required,uuid"`
	BoardID string `json:"-" validate:"required,uuid"`
	Title   string `json:"title" validate:"required"`
}

type UpdateCard struct {
	CardID      string     `json:"-" validate:"required,uuid"`
	Title       *string    `json:"title" validate:"omitempty"`
	Description *string    `json:"description" validate:"omitempty"`
	Position    *float64   `json:"position" validate:"omitempty"`
	Cover       *string    `json:"cover" validate:"omitempty"`
	Archived    *bool      `json:"archived" validate:"omitempty"`
	Completed   *bool      `json:"completed" validate:"omitempty"`
	StartDate   *time.Time `json:"startDate" validate:"omitempty"`
	DueDate     *time.Time `json:"dueDate" validate:"omitempty"`
}

type Card struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Position    float64 `json:"position"`
	Cover       string  `json:"cover,omitempty"`
	Archived    bool    `json:"archived"`
	Completed   bool    `json:"completed"`

	Members   []CardMember `json:"members,omitempty"`
	Labels    []CardLabel  `json:"labels,omitempty"`
	Checklist []Checklists `json:"checklist,omitempty"`
	// Attachments []Attachment `json:"attachments"`
}

type CardMember struct {
	UserID   string `json:"userID"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CardLabel struct {
	LabelID string `json:"labelID"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}

type Checklists struct {
	Title string          `json:"title"`
	Items []ChecklistItem `json:"items"`
}

type ChecklistItem struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	UserID string `json:"userID"`
}
