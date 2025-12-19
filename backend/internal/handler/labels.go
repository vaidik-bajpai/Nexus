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

func (h *handler) handleModifyLabel(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("update label called")
	var modifyLabel *types.ModifyLabel
	if err := helper.ReadJSON(r, &modifyLabel); err != nil {
		helper.BadRequest(h.logger, w, "invalid request body", err)
		return
	}

	if err := h.validator.Struct(modifyLabel); err != nil {
		helper.BadRequest(h.logger, w, "validation failed", err)
		return
	}

	switch modifyLabel.Type {
	case "update":
		if err := h.store.UpdateLabel(r.Context(), modifyLabel); err != nil {
			helper.InternalServerError(h.logger, w, "failed to update label", err)
			return
		}
	case "delete":
		if err := h.store.DeleteLabel(r.Context(), modifyLabel); err != nil {
			helper.InternalServerError(h.logger, w, "failed to delete label", err)
			return
		}
	default:
		helper.BadRequest(h.logger, w, "invalid request body", nil)
		return
	}

	helper.Created(h.logger, w, "label updated successfully", nil)
}

func (h *handler) handleAddLabelToCard(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("add label to card called")
	cardID := r.PathValue("cardID")
	boardID := r.PathValue("boardID")
	var addLabelToCard *types.ToggleLabelToCard
	if err := helper.ReadJSON(r, &addLabelToCard); err != nil {
		helper.BadRequest(h.logger, w, "invalid request body", err)
		return
	}

	addLabelToCard.CardID = cardID
	addLabelToCard.BoardID = boardID

	if err := h.validator.Struct(addLabelToCard); err != nil {
		helper.BadRequest(h.logger, w, "validation failed", err)
		return
	}

	switch addLabelToCard.Type {
	case "add":
		if err := h.store.AddLabelToCard(r.Context(), addLabelToCard); err != nil {
			helper.InternalServerError(h.logger, w, "failed to add label to card", err)
			return
		}
	case "remove":
		if err := h.store.RemoveLabelFromCard(r.Context(), addLabelToCard); err != nil {
			helper.InternalServerError(h.logger, w, "failed to remove label from card", err)
			return
		}
	default:
		helper.BadRequest(h.logger, w, "invalid request body", nil)
		return
	}

	helper.Created(h.logger, w, "label added to card successfully", nil)
}
