package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateCredentialsUser(ctx context.Context, user *types.CreateCredentialsUser) error {
	userID := uuid.New().String()

	userTx := s.db.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.PasswordHash),
		db.User.ID.Set(userID),
		db.User.Username.Set(user.Username),
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

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	user, err := s.db.User.FindFirst(
		db.User.Email.Equals(email),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	password, ok := user.Password()
	if !ok {
		password = ""
	}

	username, ok := user.Username()
	if !ok {
		username = ""
	}

	return &types.User{
		ID:        user.ID,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Store) UpdateUserRefreshToken(ctx context.Context, userID string, refreshToken string) error {
	_, err := s.db.User.FindUnique(
		db.User.ID.Equals(userID),
	).Update(
		db.User.RefreshToken.Set(refreshToken),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateOAuthUser(ctx context.Context, user *types.CreateOAuthUser) error {
	userID := uuid.New().String()

	userTx := s.db.User.UpsertOne(
		db.User.Email.Equals(user.Email),
	).Create(
		db.User.Email.Set(user.Email),
		db.User.ID.Set(userID),
		db.User.EmailVerified.Set(time.Now()),
	).Update(
		db.User.Email.Set(user.Email),
	).Tx()

	accTx := s.db.Account.UpsertOne(
		db.Account.ProviderProviderAccountID(
			db.Account.Provider.Equals(user.Provider),
			db.Account.ProviderAccountID.Equals(user.ProviderAccountID),
		),
	).Create(
		db.Account.Type.Set("oauth"),
		db.Account.Provider.Set(user.Provider),
		db.Account.ProviderAccountID.Set(user.ProviderAccountID),
		db.Account.User.Link(
			db.User.ID.Equals(userID),
		),
	).Update(
		db.Account.Provider.Set(user.Provider),
	).Tx()

	if err := s.db.Prisma.Transaction(userTx, accTx).Exec(ctx); err != nil {
		return err
	}

	usr := userTx.Result()
	if user == nil {
		return errors.New("user not found")
	}

	user.ID = usr.ID

	return nil
}

func (s *Store) CreateToken(ctx context.Context, token *types.Token) error {
	_, err := s.db.Token.UpsertOne(
		db.Token.UIDScope(
			db.Token.UID.Equals(token.UserID),
			db.Token.Scope.Equals(token.Scope),
		),
	).Create(
		db.Token.Token.Set(token.Token),
		db.Token.TTL.Set(token.TTL),
		db.Token.Scope.Set(token.Scope),
		db.Token.User.Link(
			db.User.ID.Equals(token.UserID),
		),
	).Update(
		db.Token.Token.Set(token.Token),
		db.Token.TTL.Set(token.TTL),
	).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByToken(ctx context.Context, token string) (*types.TokenUser, error) {
	t, err := s.db.Token.FindFirst(
		db.Token.Token.Equals(token),
	).With(
		db.Token.User.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	usr := t.User()
	if usr == nil {
		return nil, errors.New("user not found")
	}

	username, ok := usr.Username()
	if !ok {
		username = ""
	}

	password, ok := usr.Password()
	if !ok {
		password = ""
	}

	return &types.TokenUser{
		Token: types.Token{
			UserID: usr.ID,
			Token:  t.Token,
			TTL:    t.TTL,
			Scope:  t.Scope,
		},
		User: types.User{
			ID:        usr.ID,
			Username:  username,
			Email:     usr.Email,
			Password:  password,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		},
	}, nil
}

func (s *Store) UpdateUserPassword(ctx context.Context, userID string, password string) error {
	_, err := s.db.User.FindUnique(
		db.User.ID.Equals(userID),
	).Update(
		db.User.Password.Set(password),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
