package types

type CreateList struct {
	BoardID  string  `json:"-" validate:"required,uuid"`
	Name     string  `json:"name" validate:"required"`
	Position float64 `json:"position" validate:"required"`
	Color    string  `json:"color" validate:"omitempty"`
}

type UpdateList struct {
	ListID    string   `json:"-" validate:"required,uuid"`
	Name      *string  `json:"name" validate:"omitempty"`
	Position  *float64 `json:"position" validate:"omitempty"`
	Color     *string  `json:"color" validate:"omitempty"`
	Archived  *bool    `json:"archived" validate:"omitempty"`
	Collapsed *bool    `json:"collapsed" validate:"omitempty"`
}
