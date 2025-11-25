package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) GetBoardMember(ctx context.Context, boardID, memberID string) (*types.BoardMember, error) {
	boardMember, err := s.db.BoardMember.FindUnique(
		db.BoardMember.BoardIDUserID(
			db.BoardMember.BoardID.Equals(boardID),
			db.BoardMember.UserID.Equals(memberID),
		),
	).With(
		db.BoardMember.Board.Fetch().Select(
			db.Board.Name.Field(),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &types.BoardMember{
		BoardName: boardMember.Board().Name,
		BoardID:   boardID,
		UserID:    memberID,
		Role:      boardMember.Role,
	}, nil
}
