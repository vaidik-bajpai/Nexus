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
