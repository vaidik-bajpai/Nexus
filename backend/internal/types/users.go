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

type CreateOAuthUser struct {
	ID                string `json:"id"`
	Email             string `json:"email" validate:"required,email"`
	Provider          string `json:"provider" validate:"required,oneof=google"`
	ProviderAccountID string `json:"provider_account_id" validate:"required"`
}

type GoogleOAuthUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type UserContextKey string

var UserCtxKey = UserContextKey("user")
