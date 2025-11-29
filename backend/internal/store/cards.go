package store

import (
	"context"
	"log"

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
	).Exec(ctx)
	return err
}

func (s *Store) UpdateCard(ctx context.Context, cardID string, card *types.UpdateCard) error {
	log.Println(*card.Cover)
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
	).Exec(ctx)
	return err
}

func (s *Store) GetCardDetail(ctx context.Context, cardID string) (*types.Card, error) {
	card, err := s.db.Card.FindUnique(
		db.Card.ID.Equals(cardID),
	).With(
		db.Card.CardMembers.Fetch(),
		db.Card.CardLabels.Fetch().With(
			db.CardLabel.Label.Fetch(),
		),
		db.Card.Checklists.Fetch().With(
			db.Checklist.Items.Fetch(),
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

	for _, member := range card.CardMembers() {
		username, ok := member.User().Username()
		if !ok {
			username = ""
		}

		avatar, ok := member.User().Avatar()
		if !ok {
			avatar = ""
		}

		res.Members = append(res.Members, types.CardMember{
			UserID:   member.User().ID,
			Avatar:   avatar,
			Username: username,
			Email:    member.User().Email,
		})
	}

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
