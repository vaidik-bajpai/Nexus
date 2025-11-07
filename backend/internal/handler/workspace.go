package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
)

func (h *handler) handleCreateWorkspace(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling workspace creation")

	user := r.Context().Value(types.UserCtxKey).(*types.User)
	h.logger.Debug("user", zap.Any("user", user))

	var workspace types.CreateWorkspace
	if err := helper.ReadJSON(r, &workspace); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to read request body",
		})
		return
	}

	if err := h.validator.Struct(workspace); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
		})
		return
	}

	h.logger.Debug("request body validated")
	workspace.UserID = user.ID

	if err := h.store.CreateWorkspace(r.Context(), &workspace); err != nil {
		h.logger.Error("failed to create workspace", zap.Error(err))
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create workspace",
		})
		return
	}

	h.logger.Debug("workspace created successfully")

	helper.WriteJSON(w, http.StatusCreated, &types.Response{
		Status:  http.StatusCreated,
		Message: "workspace created successfully",
		Data:    nil,
	})
}

func (h *handler) handleListWorkspaces(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling workspace list")

	user := r.Context().Value(types.UserCtxKey).(*types.User)
	h.logger.Debug("user", zap.Any("user", user))

	workspaces, err := h.store.ListWorkspaces(r.Context(), user.ID)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to list workspaces",
		})
		return
	}

	h.logger.Debug("workspaces listed successfully", zap.Any("workspaces", workspaces))

	helper.WriteJSON(w, http.StatusOK, &types.Response{
		Status:  http.StatusOK,
		Message: "workspaces listed successfully",
		Data:    workspaces,
	})
}
