package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateCredentialsUser(ctx context.Context, user *types.CreateCredentialsUser) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.User), args.Error(1)
}

func (m *MockStore) UpdateUserRefreshToken(ctx context.Context, userID string, refreshToken string) error {
	args := m.Called(ctx, userID, refreshToken)
	return args.Error(0)
}

func (m *MockStore) CreateOAuthUser(ctx context.Context, user *types.CreateOAuthUser) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockStore) CreateToken(ctx context.Context, token *types.Token) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

func (m *MockStore) GetUserByToken(ctx context.Context, token string) (*types.TokenUser, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.TokenUser), args.Error(1)
}

func (m *MockStore) UpdateUserPassword(ctx context.Context, userID string, password string) error {
	args := m.Called(ctx, userID, password)
	return args.Error(0)
}

func (m *MockStore) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockStore) CreateBoard(ctx context.Context, board *types.CreateBoard) error {
	args := m.Called(ctx, board)
	return args.Error(0)
}

func (m *MockStore) ListBoards(ctx context.Context, ownerID string, paginate *types.Paginate) ([]*types.Board, error) {
	args := m.Called(ctx, ownerID, paginate)
	return nil, args.Error(0)
}

func (m *MockStore) GetBoardMember(ctx context.Context, boardID, memberID string) (*types.BoardMember, error) {
	args := m.Called(ctx, boardID, memberID)
	return nil, args.Error(0)
}

func (m *MockStore) CreateBoardInvitation(ctx context.Context, invitation *types.BoardInvitation) error {
	args := m.Called(ctx, invitation)
	return args.Error(0)
}

func (m *MockStore) IsABoardMember(ctx context.Context, email, boardID string) (bool, error) {
	args := m.Called(ctx, email, boardID)
	return true, args.Error(0)
}

func (m *MockStore) AcceptBoardInvitation(ctx context.Context, token, userID, role string) error {
	args := m.Called(ctx, token, userID, role)
	return args.Error(0)
}

func (m *MockStore) GetBoardInvitationByToken(ctx context.Context, token string) (*types.BoardInvitation, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.BoardInvitation), args.Error(1)
}

func (m *MockStore) UpdateBoard(ctx context.Context, board *types.UpdateBoard) error {
	args := m.Called(ctx, board)
	return args.Error(0)
}

func (m *MockStore) DeleteBoard(ctx context.Context, boardID string) error {
	args := m.Called(ctx, boardID)
	return args.Error(0)
}

func (m *MockStore) CreateList(ctx context.Context, list *types.CreateList) error {
	args := m.Called(ctx, list)
	return args.Error(0)
}

func (m *MockStore) UpdateList(ctx context.Context, list *types.UpdateList) error {
	args := m.Called(ctx, list)
	return args.Error(0)
}

func (m *MockStore) DeleteList(ctx context.Context, listID string) error {
	args := m.Called(ctx, listID)
	return args.Error(0)
}

func (m *MockStore) CreateCard(ctx context.Context, card *types.CreateCard) error {
	args := m.Called(ctx, card)
	return args.Error(0)
}

func (m *MockStore) UpdateCard(ctx context.Context, cardID string, card *types.UpdateCard) error {
	args := m.Called(ctx, cardID, card)
	return args.Error(0)
}

func (m *MockStore) GetCardDetail(ctx context.Context, cardID string) (*types.CompleteCard, error) {
	args := m.Called(ctx, cardID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.CompleteCard), args.Error(1)
}

func (m *MockStore) DeleteCard(ctx context.Context, cardID string) error {
	args := m.Called(ctx, cardID)
	return args.Error(0)
}

func (m *MockStore) GetCardsAndLists(ctx context.Context, boardID string) (*types.BoardDetail, error) {
	args := m.Called(ctx, boardID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.BoardDetail), args.Error(1)
}

func (m *MockStore) ToggleCardMembership(ctx context.Context, member *types.ToggleCardMembership) error {
	args := m.Called(ctx, member)
	return args.Error(0)
}

func (m *MockStore) CreateLabel(ctx context.Context, label *types.CreateLabel) error {
	args := m.Called(ctx, label)
	return args.Error(0)
}

func (m *MockStore) UpdateLabel(ctx context.Context, label *types.ModifyLabel) error {
	args := m.Called(ctx, label)
	return args.Error(0)
}

func (m *MockStore) DeleteLabel(ctx context.Context, label *types.ModifyLabel) error {
	args := m.Called(ctx, label)
	return args.Error(0)
}

func (m *MockStore) AddLabelToCard(ctx context.Context, label *types.ToggleLabelToCard) error {
	args := m.Called(ctx, label)
	return args.Error(0)
}

func (m *MockStore) RemoveLabelFromCard(ctx context.Context, label *types.ToggleLabelToCard) error {
	args := m.Called(ctx, label)
	return args.Error(0)
}

func (m *MockStore) ListBoardLabels(ctx context.Context, boardID string) ([]*types.ListLabels, error) {
	args := m.Called(ctx, boardID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*types.ListLabels), args.Error(1)
}

func (m *MockStore) ListCardLabels(ctx context.Context, boardID, cardID string) ([]*types.ListCardLabels, error) {
	args := m.Called(ctx, boardID, cardID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*types.ListCardLabels), args.Error(1)
}

func (m *MockStore) GetBoard(ctx context.Context, boardID string) (*types.CompleteBoard, error) {
	args := m.Called(ctx, boardID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.CompleteBoard), args.Error(1)
}

func (m *MockStore) AddChecklistToCard(ctx context.Context, addChecklist *types.AddChecklist) error {
	args := m.Called(ctx, addChecklist)
	return args.Error(0)
}

func (m *MockStore) GetChecklist(ctx context.Context, checklistID string) (*types.Checklist, error) {
	args := m.Called(ctx, checklistID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.Checklist), args.Error(1)
}

func (m *MockStore) DeleteChecklist(ctx context.Context, checklistID string) error {
	args := m.Called(ctx, checklistID)
	return args.Error(0)
}

func (m *MockStore) AddChecklistItem(ctx context.Context, addItem *types.AddChecklistItem) (*types.ChecklistItem, error) {
	args := m.Called(ctx, addItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.ChecklistItem), args.Error(1)
}

func (m *MockStore) DeleteChecklistItem(ctx context.Context, itemID string) error {
	args := m.Called(ctx, itemID)
	return args.Error(0)
}

func (m *MockStore) UpdateChecklistItem(ctx context.Context, updateItem *types.UpdateChecklistItem) error {
	args := m.Called(ctx, updateItem)
	return args.Error(0)
}
