package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
)

func (h *handler) handleCreateLabel(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("create label called")
	var req types.CreateLabel
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.BadRequest(h.logger, w, "invalid request body", err)
		return
	}

	h.logger.Info("payload", zap.Any("payload", req))

	req.BoardID = r.PathValue("boardID")

	if err := h.validator.Struct(req); err != nil {
		helper.BadRequest(h.logger, w, "validation failed", err)
		return
	}

	if err := h.store.CreateLabel(r.Context(), &req); err != nil {
		helper.InternalServerError(h.logger, w, "failed to create label", err)
		return
	}

	helper.Created(h.logger, w, "label created successfully", nil)
}
