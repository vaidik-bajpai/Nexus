package handler

import (
	"net/http"

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
