package types

type CreateLabel struct {
	BoardID string `json:"-" validate:"required,uuid"`
	Name    string `json:"name" validate:"required"`
	Color   string `json:"color" validate:"required,hexcolor"`
}
