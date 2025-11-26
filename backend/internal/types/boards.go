package types

import "time"

type CreateBoard struct {
	Name       string `json:"name" validate:"required,max=20"`
	Background string `json:"background" validate:"required,color_or_url"`
	Visibility string `json:"visibility,omitempty" validate:"omitempty,oneof=private team public"`

	OwnerID string `json:"-" validate:"-"`
}

type Board struct {
	Name       string `json:"name"`
	Background string `json:"background"`
}

type BoardInvitation struct {
	BoardID   string    `json:"-" validate:"required,uuid"`
	Email     string    `json:"email" validate:"required,email"`
	InvitedBy string    `json:"-" validate:"required,uuid"`
	Token     string    `json:"-" validate:"required"`
	ExpiredAt time.Time `json:"-" validate:"required,gt"`
	Role      string    `json:"role" validate:"required,oneof=observer member"`
}

type UpdateBoard struct {
	BoardID     string  `json:"-"`
	Name        *string `json:"name" validate:"omitempty,max=20"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Visibility  *string `json:"visibility" validate:"omitempty,oneof=private team public"`
	Background  *string `json:"background" validate:"omitempty,color_or_url"`
	Archived    *bool   `json:"archived" validate:"omitempty"`
}

type BoardContextKey string

var BoardCtxKey = BoardContextKey("board")
