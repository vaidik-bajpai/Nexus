package types

type CreateBoard struct {
	Name       string `json:"name" validate:"required,max=20"`
	Background string `json:"background" validate:"required,color_or_url"`
	Visibility string `json:"visibility,omitempty" validate:"omitempty,oneof=private team public"`

	OwnerID string `json:"-" validate:"-"`
}
