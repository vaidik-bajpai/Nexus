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
