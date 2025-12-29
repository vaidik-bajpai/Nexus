package store

import (
	"context"
	"time"

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
			ID:         board.ID,
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

func (s *Store) GetCardsAndLists(ctx context.Context, boardID string) (*types.BoardDetail, error) {
	dbBoard, err := s.db.Board.FindUnique(
		db.Board.ID.Equals(boardID),
	).With(
		db.Board.Lists.Fetch().Select(
			db.List.ID.Field(),
			db.List.Name.Field(),
			db.List.Position.Field(),
		).OrderBy(
			db.List.Position.Order(db.SortOrder("asc")),
		).With(
			db.List.Cards.Fetch().Select(
				db.Card.ID.Field(),
				db.Card.Title.Field(),
				db.Card.Description.Field(),
				db.Card.Completed.Field(),
				db.Card.Cover.Field(),
				db.Card.CoverSize.Field(),
				db.Card.DueDate.Field(),
				db.Card.Position.Field(),
			).With(
				db.Card.CardLabels.Fetch().With(
					db.CardLabel.Label.Fetch().Select(
						db.Label.ID.Field(),
						db.Label.Color.Field(),
						db.Label.Name.Field(),
					),
				),
				db.Card.CardMembers.Fetch().Select(
					db.CardMember.UserID.Field(),
				),
			).OrderBy(
				db.Card.Position.Order(db.SortOrder("asc")),
			),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var board types.BoardDetail
	board.ID = dbBoard.ID

	for _, list := range dbBoard.Lists() {
		board.Lists = append(board.Lists, &types.List{
			ID:       list.ID,
			Name:     list.Name,
			Position: list.Position,
		})

		cards := list.Cards()
		for _, card := range cards {
			cardCover, ok := card.Cover()
			if !ok {
				cardCover = ""
			}

			cardCoverSize, ok := card.CoverSize()
			if !ok {
				cardCoverSize = "normal"
			}

			cardDueDate, ok := card.DueDate()
			if !ok {
				cardDueDate = time.Time{}
			}

			cardDescription, ok := card.Description()
			if !ok {
				cardDescription = ""
			}

			var cardLabels []*types.BoardLabel
			for _, label := range card.CardLabels() {
				l := label.Label()
				cardLabels = append(cardLabels, &types.BoardLabel{
					LabelID: l.ID,
					Color:   l.Color,
					Name:    l.Name,
				})
			}

			var memberIDs []string
			for _, member := range card.CardMembers() {
				memberIDs = append(memberIDs, member.UserID)
			}

			board.Cards = append(board.Cards, &types.MinimalCard{
				ID:          card.ID,
				ListID:      list.ID,
				BoardID:     boardID,
				Title:       card.Title,
				Description: cardDescription,
				Cover:       cardCover,
				CoverSize:   cardCoverSize,
				Due:         cardDueDate,
				Completed:   card.Completed,
				Position:    card.Position,
				Labels:      cardLabels,
				MemberIDs:   memberIDs,
			})
		}
	}

	return &board, nil
}

func (s *Store) GetBoard(ctx context.Context, boardID string) (*types.CompleteBoard, error) {
	dbBoard, err := s.db.Board.FindUnique(
		db.Board.ID.Equals(boardID),
	).Select(
		db.Board.ID.Field(),
		db.Board.Name.Field(),
		db.Board.Background.Field(),
	).With(
		db.Board.Labels.Fetch().Select(
			db.Label.ID.Field(),
			db.Label.Color.Field(),
			db.Label.Name.Field(),
		),
		db.Board.BoardMembers.Fetch().Select(
			db.BoardMember.UserID.Field(),
			db.BoardMember.Role.Field(),
		).With(
			db.BoardMember.User.Fetch().Select(
				db.User.Email.Field(),
				db.User.FirstName.Field(),
				db.User.LastName.Field(),
			),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var board types.CompleteBoard
	board.ID = dbBoard.ID
	board.Name = dbBoard.Name
	background, ok := dbBoard.Background()
	if !ok {
		background = "#FFFFFF"
	}
	board.Background = background

	for _, label := range dbBoard.Labels() {
		board.Labels = append(board.Labels, &types.BoardLabel{
			LabelID: label.ID,
			Color:   label.Color,
			Name:    label.Name,
		})
	}

	for _, member := range dbBoard.BoardMembers() {
		firstName, ok := member.User().FirstName()
		if !ok {
			firstName = ""
		}
		lastName, ok := member.User().LastName()
		if !ok {
			lastName = ""
		}

		username, ok := member.User().Username()
		if !ok {
			username = ""
		}

		board.Members = append(board.Members, &types.BoardMembers{
			ID:       member.UserID,
			Email:    member.User().Email,
			FullName: firstName + " " + lastName,
			Username: username,
			Role:     member.Role,
		})
	}

	return &board, nil
}
