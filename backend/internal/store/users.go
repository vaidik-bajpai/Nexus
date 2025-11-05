package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateCredentialsUser(ctx context.Context, user *types.CreateCredentialsUser) error {
	userID := uuid.New().String()

	userTx := s.db.User.CreateOne(
		db.User.Username.Set(user.Username),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.PasswordHash),
		db.User.ID.Set(userID),
	).Tx()

	tokenTx := s.db.Token.CreateOne(
		db.Token.Token.Set(user.Token),
		db.Token.TTL.Set(user.TokenTTL),
		db.Token.Scope.Set(user.TokenScope),
		db.Token.User.Link(
			db.User.ID.Equals(userID),
		),
	).Tx()

	if err := s.db.Prisma.Transaction(userTx, tokenTx).Exec(ctx); err != nil {
		return err
	}

	return nil
}
