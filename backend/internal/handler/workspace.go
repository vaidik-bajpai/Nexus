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
