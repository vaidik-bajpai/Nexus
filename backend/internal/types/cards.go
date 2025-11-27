package types

type CreateCard struct {
	ListID  string `json:"-" validate:"required,uuid"`
	UserID  string `json:"-" validate:"required,uuid"`
	BoardID string `json:"-" validate:"required,uuid"`
	Title   string `json:"title" validate:"required"`
}
