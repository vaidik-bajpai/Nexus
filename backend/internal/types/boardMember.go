package types

type BoardMember struct {
	BoardName string
	BoardID   string
	UserID    string
	Role      string
	Members   []*Member
}

type Member struct {
	Name  string
	Email string
	Role  string
}
