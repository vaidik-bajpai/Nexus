package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateWorkspace(ctx context.Context, workspace *types.CreateWorkspace) error {
	workspaceID := uuid.New().String()

	workspaceTx := s.db.Workspace.CreateOne(
		db.Workspace.Name.Set(workspace.Name),
		db.Workspace.Description.Set(workspace.Description),
		db.Workspace.User.Link(
			db.User.ID.Equals(workspace.UserID),
		),
		db.Workspace.ID.Set(workspaceID),
	).Tx()

	workspaceMemberTx := s.db.WorkspaceMember.CreateOne(
		db.WorkspaceMember.Role.Set("admin"),
		db.WorkspaceMember.Workspace.Link(
			db.Workspace.ID.Equals(workspaceID),
		),
		db.WorkspaceMember.User.Link(
			db.User.ID.Equals(workspace.UserID),
		),
	).Tx()

	if err := s.db.Prisma.Transaction(workspaceTx, workspaceMemberTx).Exec(ctx); err != nil {
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
	).With(
		db.Workspace.WorkspaceMember.Fetch().With(
			db.WorkspaceMember.User.Fetch().Select(
				db.User.Email.Field(),
			),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var members []*types.WorkspaceMember
	for _, member := range workspace.WorkspaceMember() {
		joinedAt, ok := member.JoinedAt()
		if !ok {
			joinedAt = time.Time{}
		}

		members = append(members, &types.WorkspaceMember{
			ID:          member.ID,
			UserID:      member.UserID,
			Email:       member.User().Email,
			WorkspaceID: member.WorkspaceID,
			Role:        member.Role,
			JoinedAt:    joinedAt,
			InvitedAt:   member.InvitedAt,
		})
	}

	return &types.Workspace{
		ID:          workspace.ID,
		UserID:      workspace.UserID,
		Name:        workspace.Name,
		Description: workspace.Description,
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
		Members:     members,
	}, nil
}
