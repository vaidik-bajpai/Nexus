package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateBoard(ctx context.Context, board *types.CreateBoard) error {
	boardID := uuid.New().String()
	boardTxn := s.db.Board.CreateOne(
		db.Board.Name.Set(board.Name),
		db.Board.Owner.Link(
			db.User.ID.Equals(board.OwnerID),
		),
		db.Board.ID.Set(boardID),
		db.Board.Visibility.SetIfPresent(&board.Visibility),
		db.Board.Background.SetIfPresent(&board.Background),
	).Tx()

	boardMemTxn := s.db.BoardMember.CreateOne(
		db.BoardMember.Board.Link(
			db.Board.ID.Equals(boardID),
		),
		db.BoardMember.User.Link(
			db.User.ID.Equals(board.OwnerID),
		),
		db.BoardMember.Role.Set("admin"),
	).Tx()

	return s.db.Prisma.Transaction(boardTxn, boardMemTxn).Exec(ctx)
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

func (s *Store) CreateBoardInvitation(ctx context.Context, invitation *types.BoardInvitation) error {
	_, err := s.db.BoardInvitation.CreateOne(
		db.BoardInvitation.Email.Set(invitation.Email),
		db.BoardInvitation.Token.Set(invitation.Token),
		db.BoardInvitation.Role.Set(invitation.Role),
		db.BoardInvitation.ExpiresAt.Set(invitation.ExpiredAt),
		db.BoardInvitation.Board.Link(
			db.Board.ID.Equals(invitation.BoardID),
		),
		db.BoardInvitation.InvitedByUser.Link(
			db.User.Email.Equals(invitation.Email),
		),
	).Exec(ctx)
	return err
}

func (s *Store) IsABoardMember(ctx context.Context, email, boardID string) (bool, error) {
	members, err := s.db.BoardMember.FindMany(
		db.BoardMember.BoardID.Equals(boardID),
		db.BoardMember.User.Where(
			db.User.Email.Equals(email),
		),
	).Exec(ctx)
	if err != nil {
		return false, err
	}

	return len(members) == 0, nil
}

func (s *Store) AcceptBoardInvitation(ctx context.Context, token, userID, role string) error {
	// First, retrieve the invitation to get the associated board ID
	invitation, err := s.db.BoardInvitation.FindUnique(
		db.BoardInvitation.Token.Equals(token),
	).Exec(ctx)
	if err != nil {
		return err
	}

	// Update the invitation status to "accepted"
	boardTxn := s.db.BoardInvitation.FindUnique(
		db.BoardInvitation.Token.Equals(token),
	).Update(
		db.BoardInvitation.Status.Set("accepted"),
	).Tx()

	// Create a board member linking to the correct board ID from the invitation
	memberTxn := s.db.BoardMember.CreateOne(
		db.BoardMember.Board.Link(
			db.Board.ID.Equals(invitation.BoardID),
		),
		db.BoardMember.User.Link(
			db.User.ID.Equals(userID),
		),
		db.BoardMember.Role.Set(role),
	).Tx()

	// Execute both transactions atomically
	return s.db.Prisma.Transaction(boardTxn, memberTxn).Exec(ctx)
}

func (s *Store) UpdateBoard(ctx context.Context, board *types.UpdateBoard) error {
	_, err := s.db.Board.FindUnique(
		db.Board.ID.Equals(board.BoardID),
	).Update(
		db.Board.Name.SetIfPresent(board.Name),
		db.Board.Description.SetIfPresent(board.Description),
		db.Board.Visibility.SetIfPresent(board.Visibility),
		db.Board.Background.SetIfPresent(board.Background),
		db.Board.Archived.SetIfPresent(board.Archived),
	).Exec(ctx)
	return err
}

func (s *Store) DeleteBoard(ctx context.Context, boardID string) error {
	_, err := s.db.Board.FindUnique(
		db.Board.ID.Equals(boardID),
	).Delete().Exec(ctx)
	return err
}
