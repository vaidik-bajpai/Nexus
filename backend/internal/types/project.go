package types

type CreateProject struct {
	UserID      string `json:"-" validate:"-"`
	WorkspaceID string `json:"-" validate:"-"`
	Name        string `json:"name" validate:"required,max=32"`
	Description string `json:"description" validate:"required,max=1024"`
	Color       string `json:"color" validate:"required,hexcolor"`
	Status      string `json:"status" validate:"required,oneof=active archived completed"`
}
