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

func (s *Store) DeleteChecklist(ctx context.Context, checklistID string) error {
	_, err := s.db.Checklist.FindUnique(
		db.Checklist.ID.Equals(checklistID),
	).Delete().Exec(ctx)
	return err
}

func (s *Store) AddChecklistItem(ctx context.Context, addItem *types.AddChecklistItem) (*types.ChecklistItem, error) {
	item, err := s.db.ChecklistItem.CreateOne(
		db.ChecklistItem.Text.Set(addItem.Name),
		db.ChecklistItem.Checklist.Link(
			db.Checklist.ID.Equals(addItem.ChecklistID),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &types.ChecklistItem{
		ID:        item.ID,
		Name:      item.Text,
		Completed: item.Completed,
		Position:  item.Position,
	}, nil
}

func (s *Store) DeleteChecklistItem(ctx context.Context, itemID string) error {
	_, err := s.db.ChecklistItem.FindUnique(
		db.ChecklistItem.ID.Equals(itemID),
	).Delete().Exec(ctx)
	return err
}

func (s *Store) UpdateChecklistItem(ctx context.Context, updateItem *types.UpdateChecklistItem) error {
	_, err := s.db.ChecklistItem.FindUnique(
		db.ChecklistItem.ID.Equals(updateItem.ItemID),
	).Update(
		db.ChecklistItem.Text.SetIfPresent(updateItem.Name),
		db.ChecklistItem.Completed.SetIfPresent(updateItem.Completed),
	).Exec(ctx)
	return err
}
