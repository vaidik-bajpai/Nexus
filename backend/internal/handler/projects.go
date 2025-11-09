package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
)

func (h *handler) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling project creation")

	user := r.Context().Value(types.UserCtxKey).(*types.User)
	h.logger.Debug("user", zap.Any("user", user))

	workspaceID := r.PathValue("workspace_id")
	if workspaceID == "" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "workspace id is required",
		})
		return
	}

	var project types.CreateProject
	if err := helper.ReadJSON(r, &project); err != nil {
		h.logger.Error("failed to read request body", zap.Error(err))
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
		})
		return
	}

	if err := h.validator.Struct(project); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
		})
		return
	}

	project.UserID = user.ID
	project.WorkspaceID = workspaceID

	if err := h.store.CreateProject(r.Context(), &project); err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create project",
		})
		return
	}

	helper.WriteJSON(w, http.StatusCreated, &types.Response{
		Status:  http.StatusCreated,
		Message: "project created successfully",
		Data:    nil,
	})
}

func (h *handler) handleListProjects(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling project listing")

	user := r.Context().Value(types.UserCtxKey).(*types.User)
	h.logger.Debug("user", zap.Any("user", user))

	workspaceID := r.PathValue("workspace_id")
	if workspaceID == "" {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "workspace id is required",
		})
		return
	}

	projects, err := h.store.ListProjects(r.Context(), workspaceID)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to list projects",
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "projects listed successfully",
		Data:    projects,
	})
}
