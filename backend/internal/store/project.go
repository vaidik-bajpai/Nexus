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

func (s *Store) ListProjects(ctx context.Context, workspaceID string) ([]*types.Project, error) {
	projects, err := s.db.Project.FindMany(
		db.Project.WorkspaceID.Equals(workspaceID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var projectsList []*types.Project
	for _, project := range projects {
		color, ok := project.Color()
		if !ok {
			color = ""
		}

		description, ok := project.Description()
		if !ok {
			description = ""
		}

		projectsList = append(projectsList, &types.Project{
			ID:          project.ID,
			Name:        project.Name,
			Description: description,
			Status:      project.Status,
			Color:       color,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
			CreatedBy:   project.CreatedBy,
			WorkspaceID: project.WorkspaceID,
		})
	}

	return projectsList, nil
}
