package store

import (
	"context"
	"time"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateCard(ctx context.Context, card *types.CreateCard) error {
	_, err := s.db.Card.CreateOne(
		db.Card.Title.Set(card.Title),
		db.Card.List.Link(
			db.List.ID.Equals(card.ListID),
		),
		db.Card.Creator.Link(
			db.User.ID.Equals(card.UserID),
		),
		db.Card.Board.Link(
			db.Board.ID.Equals(card.BoardID),
		),
		db.Card.Position.Set(card.Position),
	).Exec(ctx)
	return err
}

func (s *Store) UpdateCard(ctx context.Context, cardID string, card *types.UpdateCard) error {
	_, err := s.db.Card.FindUnique(
		db.Card.ID.Equals(cardID),
	).Update(
		db.Card.Title.SetIfPresent(card.Title),
		db.Card.Description.SetIfPresent(card.Description),
		db.Card.Position.SetIfPresent(card.Position),
		db.Card.Cover.SetIfPresent(card.Cover),
		db.Card.CoverSize.SetIfPresent(card.CoverSize),
		db.Card.Archived.SetIfPresent(card.Archived),
		db.Card.Completed.SetIfPresent(card.Completed),
		db.Card.StartDate.SetIfPresent(card.StartDate),
		db.Card.DueDate.SetIfPresent(card.DueDate),
		db.Card.List.Link(
			db.List.ID.Equals(card.ListID),
		),
	).Exec(ctx)
	return err
}

func (s *Store) GetCardDetail(ctx context.Context, cardID string) (*types.CompleteCard, error) {
	dbCard, err := s.db.Card.FindUnique(
		db.Card.ID.Equals(cardID),
	).With(
		db.Card.CardLabels.Fetch().With(
			db.CardLabel.Label.Fetch(),
		),
		db.Card.CardMembers.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var card types.CompleteCard
	card.ID = dbCard.ID
	card.Name = dbCard.Title
	if desc, ok := dbCard.Description(); !ok {
		card.Description = ""
	} else {
		card.Description = desc
	}
	card.Position = dbCard.Position
	if cover, ok := dbCard.Cover(); !ok {
		card.Cover = ""
	} else {
		card.Cover = cover
	}
	if coverSize, ok := dbCard.CoverSize(); !ok {
		card.CoverSize = "normal"
	} else {
		card.CoverSize = coverSize
	}
	card.Archived = dbCard.Archived
	card.Completed = dbCard.Completed
	if startDate, ok := dbCard.StartDate(); !ok {
		card.Start = time.Time{}
	} else {
		card.Start = startDate
	}
	if dueDate, ok := dbCard.DueDate(); !ok {
		card.Due = time.Time{}
	} else {
		card.Due = dueDate
	}

	for _, member := range dbCard.CardMembers() {
		card.MemberIDs = append(card.MemberIDs, &member.UserID)
	}
	for _, label := range dbCard.CardLabels() {
		card.Labels = append(card.Labels, &types.BoardLabel{
			LabelID: label.LabelID,
			Name:    label.Label().Name,
			Color:   label.Label().Color,
		})
	}
	return &card, nil
}

func (s *Store) DeleteCard(ctx context.Context, cardID string) error {
	_, err := s.db.Card.FindUnique(
		db.Card.ID.Equals(cardID),
	).Delete().Exec(ctx)
	return err
}

func (s *Store) ToggleCardMembership(ctx context.Context, member *types.ToggleCardMembership) error {
	cardMember, err := s.db.CardMember.FindFirst(
		db.CardMember.CardID.Equals(member.CardID),
		db.CardMember.UserID.Equals(member.UserID),
	).Exec(ctx)
	if err != nil {
		if db.IsErrNotFound(err) {
			_, err = s.db.CardMember.CreateOne(
				db.CardMember.Card.Link(
					db.Card.ID.Equals(member.CardID),
				),
				db.CardMember.User.Link(
					db.User.ID.Equals(member.UserID),
				),
			).Exec(ctx)
			return err
		}
		return err
	}
	_, err = s.db.CardMember.FindUnique(
		db.CardMember.ID.Equals(cardMember.ID),
	).Delete().Exec(ctx)
	return err
}
