package handler

import (
	"net/http"
	"strconv"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
)

func (h *handler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling task creation")

	user := r.Context().Value(types.UserCtxKey).(*types.User)
	h.logger.Debug("user", zap.Any("user", user))

	projectID := r.PathValue("project_id")
	if projectID == "" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "project id is required",
		})
		return
	}

	var task types.CreateTask
	if err := helper.ReadJSON(r, &task); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
		})
		return
	}

	task.ProjectID = projectID
	task.CreatedBy = user.ID
	task.Status = "todo"

	if err := h.validator.Struct(task); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
		})
		return
	}

	if err := h.store.CreateTask(r.Context(), &task); err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create task",
		})
		return
	}

	helper.WriteJSON(w, http.StatusCreated, &types.Response{
		Status:  http.StatusCreated,
		Message: "task created successfully",
		Data:    nil,
	})
}

func (h *handler) handleListTasks(w http.ResponseWriter, r *http.Request) {
	user := helper.GetUserFromRequestContext(r)
	h.logger.Debug("user", zap.Any("user", user))

	project := helper.GetProjectFromRequestContext(r)
	h.logger.Debug("project", zap.Any("project", project))

	listTasks := &types.ListTasks{
		ProjectID: project.ID,
		Filter:    "created_at",
		Direction: "asc",
	}

	if filter := r.URL.Query().Get("filter"); filter != "" {
		listTasks.Filter = filter
	}

	if direction := r.URL.Query().Get("direction"); direction != "" {
		listTasks.Direction = direction
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to convert page to int",
		})
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to convert limit to int",
		})
		return
	}

	listTasks.Page = page
	listTasks.Limit = limit

	if err := h.validator.Struct(listTasks); err != nil {
		h.logger.Error("failed to validate request body", zap.Error(err))
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
		})
		return
	}

	tasks, err := h.store.ListTasks(r.Context(), listTasks)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to list tasks",
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "tasks listed successfully",
		Data:    tasks,
	})
}
