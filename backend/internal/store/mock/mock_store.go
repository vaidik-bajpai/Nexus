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
