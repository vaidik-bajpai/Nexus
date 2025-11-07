package types

type CreateWorkspace struct {
	UserID      string `json:"-"`
	Name        string `json:"name" validate:"required,max=32"`
	Description string `json:"description" validate:"required,max=1024"`
}
