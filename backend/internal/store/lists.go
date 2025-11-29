package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateList(ctx context.Context, list *types.CreateList) error {
	_, err := s.db.List.CreateOne(
		db.List.Name.Set(list.Name),
		db.List.Board.Link(
			db.Board.ID.Equals(list.BoardID),
		),
		db.List.Position.SetIfPresent(&list.Position),
		db.List.Color.SetIfPresent(&list.Color),
	).Exec(ctx)
	return err
}

func (s *Store) UpdateList(ctx context.Context, list *types.UpdateList) error {
	_, err := s.db.List.FindUnique(
		db.List.ID.Equals(list.ListID),
	).Update(
		db.List.Name.SetIfPresent(list.Name),
		db.List.Position.SetIfPresent(list.Position),
		db.List.Color.SetIfPresent(list.Color),
		db.List.Archived.SetIfPresent(list.Archived),
		db.List.Collapsed.SetIfPresent(list.Collapsed),
	).Exec(ctx)
	return err
}

func (s *Store) DeleteList(ctx context.Context, listID string) error {
	_, err := s.db.List.FindUnique(
		db.List.ID.Equals(listID),
	).Delete().Exec(ctx)
	return err
}
