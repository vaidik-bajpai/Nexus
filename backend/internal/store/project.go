package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateProject(ctx context.Context, project *types.CreateProject) error {
	_, err := s.db.Project.CreateOne(
		db.Project.Name.Set(project.Name),
		db.Project.Workspace.Link(
			db.Workspace.ID.Equals(project.WorkspaceID),
		),
		db.Project.Creator.Link(
			db.User.ID.Equals(project.UserID),
		),
		db.Project.Description.Set(project.Description),
		db.Project.Status.Set(project.Status),
		db.Project.Color.Set(project.Color),
	).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
