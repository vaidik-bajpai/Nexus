package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (s *Store) CreateTask(ctx context.Context, task *types.CreateTask) error {
	taskID := uuid.New().String()
	var txns []db.PrismaTransaction

	taskTx := s.db.Task.CreateOne(
		db.Task.Title.Set(task.Title),
		db.Task.Project.Link(
			db.Project.ID.Equals(task.ProjectID),
		),
		db.Task.Creator.Link(
			db.User.ID.Equals(task.CreatedBy),
		),
		db.Task.Description.Set(task.Description),
		db.Task.Status.Set(task.Status),
		db.Task.Priority.Set(task.Priority),
		db.Task.DueDate.Set(task.DueDate),
		db.Task.AssignedTo.SetIfPresent(task.AssignedTo),
		db.Task.ID.Set(taskID),
	).Tx()

	txns = append(txns, taskTx)

	for _, dependsOn := range task.DependsOn {
		txns = append(txns, s.db.TaskDependency.CreateOne(
			db.TaskDependency.Task.Link(
				db.Task.ID.Equals(taskID),
			),
			db.TaskDependency.DependsOn.Link(
				db.Task.ID.Equals(dependsOn),
			),
		).Tx())
	}

	if err := s.db.Prisma.Transaction(txns...).Exec(ctx); err != nil {
		return err
	}

	return nil
}
