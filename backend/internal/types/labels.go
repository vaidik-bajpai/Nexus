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

	BoardID string `json:"-" validate:"required,uuid"`
}

type ToggleLabelToCard struct {
	Type    string `json:"type" validate:"required,one of=add,remove"`
	LabelID string `json:"label_id" validate:"required,uuid"`
	CardID  string `json:"-" validate:"required,uuid"`

	BoardID string `json:"-" validate:"required,uuid"`
}
