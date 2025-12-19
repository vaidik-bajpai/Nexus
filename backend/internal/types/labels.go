package types

type CreateLabel struct {
	BoardID string `json:"-" validate:"required,uuid"`
	Name    string `json:"name" validate:"required"`
	Color   string `json:"color" validate:"required,hexcolor"`
}

type ModifyLabel struct {
	ID    string `json:"id" validate:"required,uuid"`
	Type  string `json:"type" validate:"required,one of=update,delete"`
	Name  string `json:"name" validate:"omitempty"`
	Color string `json:"color" validate:"omitempty,hexcolor"`
}
