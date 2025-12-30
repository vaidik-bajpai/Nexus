package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateLabel(ctx context.Context, label *types.CreateLabel) (*types.ListLabels, error) {
	newLabel, err := s.db.Label.CreateOne(
		db.Label.Name.Set(label.Name),
		db.Label.Color.Set(label.Color),
		db.Label.Board.Link(
			db.Board.ID.Equals(label.BoardID),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &types.ListLabels{
		ID:    newLabel.ID,
		Name:  newLabel.Name,
		Color: newLabel.Color,
	}, nil
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
	_, err := s.db.CardLabel.FindMany(
		db.CardLabel.CardID.Equals(label.CardID),
		db.CardLabel.LabelID.Equals(label.LabelID),
	).Delete().Exec(ctx)
	return err
}

func (s *Store) ListBoardLabels(ctx context.Context, boardID string) ([]*types.ListLabels, error) {
	labels, err := s.db.Label.FindMany(
		db.Label.Board.Link(
			db.Board.ID.Equals(boardID),
		),
	).Select(
		db.Label.ID.Field(),
		db.Label.Name.Field(),
		db.Label.Color.Field(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var res []*types.ListLabels
	for _, label := range labels {
		res = append(res, &types.ListLabels{
			ID:    label.ID,
			Name:  label.Name,
			Color: label.Color,
		})
	}
	return res, nil
}

func (s *Store) ListCardLabels(ctx context.Context, boardID, cardID string) ([]*types.ListCardLabels, error) {
	labels, err := s.db.Label.FindMany(
		db.Label.BoardID.Equals(boardID),
	).With(
		db.Label.CardLabels.Fetch(
			db.CardLabel.CardID.Equals(cardID),
		).Select(
			db.CardLabel.ID.Field(),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var res []*types.ListCardLabels
	for _, label := range labels {
		cardLabels := label.CardLabels()
		isChecked := len(cardLabels) > 0
		cardLabelID := ""
		if isChecked {
			cardLabelID = cardLabels[0].ID
		}
		res = append(res, &types.ListCardLabels{
			ID:        cardLabelID,
			LabelID:   label.ID,
			CardID:    cardID,
			Name:      label.Name,
			Color:     label.Color,
			IsChecked: isChecked,
		})
	}
	return res, nil
}
