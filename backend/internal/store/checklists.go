package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) AddChecklistToCard(ctx context.Context, addChecklist *types.AddChecklist) error {
	_, err := s.db.Checklist.CreateOne(
		db.Checklist.Name.Set(addChecklist.Name),
		db.Checklist.Card.Link(
			db.Card.ID.Equals(addChecklist.CardID),
		),
	).Exec(ctx)
	return err
}

func (s *Store) GetChecklist(ctx context.Context, checklistID string) (*types.Checklist, error) {
	dbChecklist, err := s.db.Checklist.FindUnique(
		db.Checklist.ID.Equals(checklistID),
	).With(
		db.Checklist.Items.Fetch().Select(
			db.ChecklistItem.ID.Field(),
			db.ChecklistItem.Text.Field(),
			db.ChecklistItem.Completed.Field(),
			db.ChecklistItem.Position.Field(),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var checklist types.Checklist
	checklist.ID = dbChecklist.ID
	checklist.Name = dbChecklist.Name
	checklist.Position = dbChecklist.Position

	checkItems := make([]*types.ChecklistItem, 0)
	for _, item := range dbChecklist.Items() {
		checkItems = append(checkItems, &types.ChecklistItem{
			ID:        item.ID,
			Name:      item.Text,
			Position:  item.Position,
			Completed: item.Completed,
		})
	}

	checklist.CheckItems = checkItems
	return &checklist, nil
}
