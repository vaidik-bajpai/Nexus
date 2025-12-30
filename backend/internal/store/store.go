package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

type Storer interface {
	CreateCredentialsUser(ctx context.Context, user *types.CreateCredentialsUser) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	UpdateUserRefreshToken(ctx context.Context, userID string, refreshToken string) error
	CreateOAuthUser(ctx context.Context, user *types.CreateOAuthUser) error
	CreateToken(ctx context.Context, token *types.Token) error
	GetUserByToken(ctx context.Context, token string) (*types.TokenUser, error)
	UpdateUserPassword(ctx context.Context, userID string, password string) error
	Close() error

	CreateBoard(ctx context.Context, board *types.CreateBoard) error
	ListBoards(ctx context.Context, ownerID string, paginate *types.Paginate) ([]*types.Board, error)
	GetBoardMember(ctx context.Context, boardID, memberID string) (*types.BoardMember, error)
	CreateBoardInvitation(ctx context.Context, invitation *types.BoardInvitation) error
	IsABoardMember(ctx context.Context, email, boardID string) (bool, error)
	AcceptBoardInvitation(ctx context.Context, token, userID, role string) error
	GetBoardInvitationByToken(ctx context.Context, token string) (*types.BoardInvitation, error)
	UpdateBoard(ctx context.Context, board *types.UpdateBoard) error
	DeleteBoard(ctx context.Context, boardID string) error
	GetCardsAndLists(ctx context.Context, boardID string) (*types.BoardDetail, error)
	GetBoard(ctx context.Context, boardID string) (*types.CompleteBoard, error)

	CreateList(ctx context.Context, list *types.CreateList) error
	UpdateList(ctx context.Context, payload *types.UpdateList) error
	DeleteList(ctx context.Context, listID string) error

	CreateCard(ctx context.Context, card *types.CreateCard) error
	UpdateCard(ctx context.Context, cardID string, card *types.UpdateCard) error
	GetCardDetail(ctx context.Context, cardID string) (*types.CompleteCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	ToggleCardMembership(ctx context.Context, member *types.ToggleCardMembership) error

	CreateLabel(ctx context.Context, label *types.CreateLabel) error
	UpdateLabel(ctx context.Context, label *types.ModifyLabel) error
	DeleteLabel(ctx context.Context, label *types.ModifyLabel) error
	AddLabelToCard(ctx context.Context, label *types.ToggleLabelToCard) error
	RemoveLabelFromCard(ctx context.Context, label *types.ToggleLabelToCard) error
	ListBoardLabels(ctx context.Context, boardID string) ([]*types.ListLabels, error)
	ListCardLabels(ctx context.Context, boardID, cardID string) ([]*types.ListCardLabels, error)
	AddChecklistToCard(ctx context.Context, addChecklist *types.AddChecklist) error
	GetChecklist(ctx context.Context, checklistID string) (*types.Checklist, error)
}

type Store struct {
	db *db.PrismaClient
}

func NewStore(db *db.PrismaClient) *Store {
	return &Store{db: db}
}

func (s *Store) Close() error {
	return s.db.Disconnect()
}
