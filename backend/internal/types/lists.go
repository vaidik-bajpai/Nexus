package types

type CreateList struct {
	BoardID  string `json:"-" validate:"required,uuid"`
	Name     string `json:"name" validate:"required"`
	Position int    `json:"position" validate:"required"`
	Color    string `json:"color" validate:"omitempty"`
}
