package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateBoard(ctx context.Context, board *types.CreateBoard) error {
	_, err := s.db.Board.CreateOne(
		db.Board.Name.Set(board.Name),
		db.Board.Owner.Link(
			db.User.ID.Equals(board.OwnerID),
		),
		db.Board.Visibility.SetIfPresent(&board.Visibility),
	).Exec(ctx)
	return err
}
