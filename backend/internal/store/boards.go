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

func (s *Store) ListBoards(ctx context.Context, ownerID string, paginate *types.Paginate) ([]*types.Board, error) {
	query := s.db.Board.FindMany(
		db.Board.UserID.Equals(ownerID),
	).Skip(paginate.Offset).Take(paginate.Size)

	if paginate.SortBy == "created_at" {
		query.OrderBy(
			db.Board.CreatedAt.Order(db.SortOrder(paginate.SortOrder)),
		)
	}

	boards, err := query.Exec(ctx)
	if err != nil {
		return nil, err
	}

	var listRes []*types.Board
	for _, board := range boards {
		background, ok := board.Background()
		if !ok {
			background = "#FFFFFF"
		}

		listRes = append(listRes, &types.Board{
			Name:       board.Name,
			Background: background,
		})
	}

	return listRes, nil
}
