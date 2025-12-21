package types

import "time"

type CreateBoard struct {
	Name       string `json:"name" validate:"required,max=20"`
	Background string `json:"background" validate:"required,color_or_url"`
	Visibility string `json:"visibility,omitempty" validate:"omitempty,oneof=private team public"`

	OwnerID string `json:"-" validate:"-"`
}

type Board struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Background string `json:"background"`
}

type BoardDetail struct {
	ID    string         `json:"id"`
	Lists []*List        `json:"lists"`
	Cards []*MinimalCard `json:"cards"`
}

type List struct {
	ID       string  `json:"id"`
	BoardID  string  `json:"board_id"`
	Name     string  `json:"name"`
	Position float64 `json:"position"`
}

type MinimalCard struct {
	ID        string        `json:"id"`
	BoardID   string        `json:"board_id"`
	ListID    string        `json:"list_id"`
	Title     string        `json:"title"`
	Cover     string        `json:"cover"`
	CoverSize string        `json:"cover_size"`
	Due       time.Time     `json:"due"`
	Completed bool          `json:"completed"`
	Position  float64       `json:"position"`
	LabelIDs  []string      `json:"label_ids"`
	MemberIDs []string      `json:"member_ids"`
	Labels    []*BoardLabel `json:"labels"`
	// Checklists
	// Attachments
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

type GetBoardDetail struct {
	BoardID string `json:"boardID" validate:"required,uuid"`
}

type CompleteBoard struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Background string          `json:"background"`
	Labels     []*BoardLabel   `json:"labels"`
	Members    []*BoardMembers `json:"members"`
}

type BoardLabel struct {
	LabelID string `json:"id"`
	Color   string `json:"color"`
	BoardID string `json:"board_id"`
	Name    string `json:"name"`
}

type BoardMembers struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type BoardContextKey string

var BoardCtxKey = BoardContextKey("board")
