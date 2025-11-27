package types

type CreateCard struct {
	ListID  string `json:"-" validate:"required,uuid"`
	UserID  string `json:"-" validate:"required,uuid"`
	BoardID string `json:"-" validate:"required,uuid"`
	Title   string `json:"title" validate:"required"`
}

type UpdateCard struct {
	CardID      string   `json:"-" validate:"required,uuid"`
	Title       *string  `json:"title" validate:"omitempty"`
	Description *string  `json:"description" validate:"omitempty"`
	Position    *float64 `json:"position" validate:"omitempty"`
	Cover       *string  `json:"cover" validate:"omitempty"`
	Archived    *bool    `json:"archived" validate:"omitempty"`
	Completed   *bool    `json:"completed" validate:"omitempty"`
}
