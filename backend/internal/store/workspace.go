package store

import (
	"context"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateWorkspace(ctx context.Context, workspace *types.CreateWorkspace) error {
	_, err := s.db.Workspace.CreateOne(
		db.Workspace.Name.Set(workspace.Name),
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

func (s *Store) ListWorkspaces(ctx context.Context, userID string) ([]*types.Workspace, error) {
	workspaces, err := s.db.Workspace.FindMany(
		db.Workspace.UserID.Equals(userID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var workspacesList []*types.Workspace
	for _, workspace := range workspaces {
		workspacesList = append(workspacesList, &types.Workspace{
			ID:          workspace.ID,
			UserID:      workspace.UserID,
			Name:        workspace.Name,
			Description: workspace.Description,
			CreatedAt:   workspace.CreatedAt,
			UpdatedAt:   workspace.UpdatedAt,
		})
	}

	return workspacesList, nil
}

func (s *Store) GetWorkspace(ctx context.Context, id string) (*types.Workspace, error) {
	workspace, err := s.db.Workspace.FindUnique(
		db.Workspace.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &types.Workspace{
		ID:          workspace.ID,
		UserID:      workspace.UserID,
		Name:        workspace.Name,
		Description: workspace.Description,
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
	}, nil
}
