package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateCard(ctx context.Context, card *types.CreateCard) error {
	_, err := s.db.Card.CreateOne(
		db.Card.Title.Set(card.Title),
		db.Card.List.Link(
			db.List.ID.Equals(card.ListID),
		),
		db.Card.Creator.Link(
			db.User.ID.Equals(card.UserID),
		),
		db.Card.Board.Link(
			db.Board.ID.Equals(card.BoardID),
		),
	).Exec(ctx)
	return err
}
