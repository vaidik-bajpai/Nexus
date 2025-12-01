package store

import (
	"context"

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

func (s *Store) GetCardDetail(ctx context.Context, cardID string) (*types.Card, error) {
	card, err := s.db.Card.FindUnique(
		db.Card.ID.Equals(cardID),
	).With(
		db.Card.CardMembers.Fetch().With(
			db.CardMember.User.Fetch(),
		),
		db.Card.CardLabels.Fetch().With(
			db.CardLabel.Label.Fetch(),
		),
		db.Card.Checklists.Fetch().With(
			db.Checklist.Items.Fetch(),
		),
		db.Card.Board.Fetch().With(
			db.Board.BoardMembers.Fetch().With(
				db.BoardMember.User.Fetch(),
			),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var res types.Card
	var ok bool
	res.ID = card.ID
	res.Title = card.Title
	res.Position = card.Position
	res.Archived = card.Archived
	res.Completed = card.Completed
	res.Description, ok = card.Description()
	if !ok {
		res.Description = ""
	}
	res.Cover, ok = card.Cover()
	if !ok {
		res.Cover = ""
	}

	cardMemberIDs := make(map[string]bool)
	for _, cm := range card.CardMembers() {
		cardMemberIDs[cm.UserID] = true
	}

	var members []types.CardMember
	for _, member := range card.Board().BoardMembers() {
		user := member.User()
		username, ok := user.Username()
		if !ok {
			username = ""
		}

		avatar, ok := user.Avatar()
		if !ok {
			avatar = ""
		}
		members = append(members, types.CardMember{
			UserID:       user.ID,
			Avatar:       avatar,
			Username:     username,
			Email:        user.Email,
			IsCardMember: cardMemberIDs[user.ID],
		})
	}

	res.Members = members

	var labels []types.CardLabel
	for _, label := range card.CardLabels() {
		labels = append(labels, types.CardLabel{
			LabelID: label.LabelID,
			Name:    label.Label().Name,
			Color:   label.Label().Color,
		})
	}

	var checklist []types.Checklists
	for _, item := range card.Checklists() {
		var listItems []types.ChecklistItem
		for _, checklistItem := range item.Items() {
			listItems = append(listItems, types.ChecklistItem{
				ID:    checklistItem.ID,
				Title: checklistItem.Text,
				Done:  checklistItem.Completed,
			})
		}
		checklist = append(checklist, types.Checklists{
			Title: item.Name,
			Items: listItems,
		})
	}

	res.Labels = labels
	res.Checklist = checklist
	return &res, nil
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
