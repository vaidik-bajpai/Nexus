package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) GetBoardInvitationByToken(ctx context.Context, token string) (*types.BoardInvitation, error) {
	invitation, err := s.db.BoardInvitation.FindUnique(
		db.BoardInvitation.Token.Equals(token),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &types.BoardInvitation{
		BoardID:   invitation.BoardID,
		Email:     invitation.Email,
		Token:     invitation.Token,
		Role:      invitation.Role,
		ExpiredAt: invitation.ExpiresAt,
	}, nil
}
