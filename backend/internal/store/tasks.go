package store

import (
	"context"
	"time"

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

func (s *Store) ListTasks(ctx context.Context, listTasks *types.ListTasks) ([]*types.Task, error) {
	query := s.db.Task.FindMany(
		db.Task.ProjectID.Equals(listTasks.ProjectID),
	)

	if listTasks.Filter == "created_at" {
		query = query.OrderBy(db.Task.CreatedAt.Order(db.SortOrder(listTasks.Direction)))
	}

	if listTasks.Filter == "priority" {
		query = query.OrderBy(db.Task.Priority.Order(db.SortOrder(listTasks.Direction)))
	}

	if listTasks.Filter == "due_date" {
		query = query.OrderBy(db.Task.DueDate.Order(db.SortOrder(listTasks.Direction)))
	}

	tasks, err := query.Skip((listTasks.Page - 1) * listTasks.Limit).Take(listTasks.Limit).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var tasksList []*types.Task
	for _, task := range tasks {
		description, ok := task.Description()
		if !ok {
			description = ""
		}

		dueDate, ok := task.DueDate()
		if !ok {
			dueDate = time.Time{}
		}

		assignedTo, ok := task.AssignedTo()
		if !ok {
			assignedTo = ""
		}

		tasksList = append(tasksList, &types.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: description,
			Status:      task.Status,
			Priority:    task.Priority,
			DueDate:     dueDate,
			AssignedTo:  assignedTo,
			CreatedBy:   task.CreatedBy,
			CreatedAt:   task.CreatedAt,
		})
	}
	return tasksList, nil
}
