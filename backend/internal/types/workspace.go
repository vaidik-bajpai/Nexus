package types

import "time"

type CreateWorkspace struct {
	UserID      string `json:"-"`
	Name        string `json:"name" validate:"required,max=32"`
	Description string `json:"description" validate:"required,max=1024"`
}

type Workspace struct {
	ID          string             `json:"id"`
	UserID      string             `json:"user_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Members     []*WorkspaceMember `json:"members"`
}

type WorkspaceMember struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Email       string    `json:"email"`
	WorkspaceID string    `json:"workspace_id"`
	Role        string    `json:"role"`
	JoinedAt    time.Time `json:"joined_at"`
	InvitedAt   time.Time `json:"invited_at"`
}
