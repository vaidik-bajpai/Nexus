package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateWorkspace(ctx context.Context, workspace *types.CreateWorkspace) error {
	_, err := s.db.Workspace.CreateOne(
		db.Workspace.Name.Equals(workspace.Name),
		db.Workspace.Description.Set(workspace.Description),
		db.Workspace.User.Link(
			db.User.ID.Equals(workspace.UserID),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
