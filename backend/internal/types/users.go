package types

import "time"

type CreateCredentialsUser struct {
	Username     string    `json:"username" validate:"required,min=3,max=32"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,min=8,max=72"`
	PasswordHash string    `json:"-" validate:"-"`
	Token        string    `json:"-" validate:"-"`
	TokenScope   string    `json:"-" validate:"-"`
	TokenTTL     time.Time `json:"-" validate:"-"`
}

type LoginCredentialsUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
