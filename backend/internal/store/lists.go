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
