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

func (s *Store) UpdateCard(ctx context.Context, cardID string, card *types.UpdateCard) error {
	_, err := s.db.Card.FindUnique(
		db.Card.ID.Equals(cardID),
	).Update(
		db.Card.Title.SetIfPresent(card.Title),
		db.Card.Description.SetIfPresent(card.Description),
		db.Card.Position.SetIfPresent(card.Position),
		db.Card.Cover.SetIfPresent(card.Cover),
		db.Card.Archived.SetIfPresent(card.Archived),
		db.Card.Completed.SetIfPresent(card.Completed),
	).Exec(ctx)
	return err
}
