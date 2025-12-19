package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateLabel(ctx context.Context, label *types.CreateLabel) error {
	_, err := s.db.Label.CreateOne(
		db.Label.Name.Set(label.Name),
		db.Label.Color.Set(label.Color),
		db.Label.Board.Link(
			db.Board.ID.Equals(label.BoardID),
		),
	).Exec(ctx)
	return err
}

func (s *Store) UpdateLabel(ctx context.Context, label *types.ModifyLabel) error {
	_, err := s.db.Label.FindUnique(
		db.Label.ID.Equals(label.ID),
	).Update(
		db.Label.Name.Set(label.Name),
		db.Label.Color.Set(label.Color),
	).Exec(ctx)
	return err
}

func (s *Store) DeleteLabel(ctx context.Context, label *types.ModifyLabel) error {
	_, err := s.db.Label.FindUnique(
		db.Label.ID.Equals(label.ID),
	).Delete().Exec(ctx)
	return err
}

func (s *Store) AddLabelToCard(ctx context.Context, label *types.ToggleLabelToCard) error {
	_, err := s.db.CardLabel.CreateOne(
		db.CardLabel.Card.Link(
			db.Card.ID.Equals(label.CardID),
		),
		db.CardLabel.Label.Link(
			db.Label.ID.Equals(label.LabelID),
		),
	).Exec(ctx)
	return err
}

func (s *Store) RemoveLabelFromCard(ctx context.Context, label *types.ToggleLabelToCard) error {
	_, err := s.db.CardLabel.FindUnique(
		db.CardLabel.ID.Equals(label.LabelID),
	).Delete().Exec(ctx)
	return err
}
