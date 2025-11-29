package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
	"go.uber.org/zap"
)

func (h *handler) handleCreateCard(w http.ResponseWriter, r *http.Request) {
	listID := r.PathValue("listID")
	boardID := r.PathValue("boardID")
	user := helper.GetUserFromRequestContext(r)
	var payload types.CreateCard
	if err := helper.ReadJSON(r, &payload); err != nil {
		helper.BadRequest(h.logger, w, "failed to read the request payload", err)
		return
	}

	payload.ListID = listID
	payload.BoardID = boardID
	payload.UserID = user.ID

	if err := h.validator.Struct(payload); err != nil {
		helper.BadRequest(h.logger, w, "failed validation on the request payload", err)
		return
	}

	if err := h.store.CreateCard(r.Context(), &payload); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "card created successfully", nil)
}

func (h *handler) handleUpdateCard(w http.ResponseWriter, r *http.Request) {
	listID := r.PathValue("listID")
	cardID := r.PathValue("cardID")
	var payload types.UpdateCard
	if err := helper.ReadJSON(r, &payload); err != nil {
		helper.BadRequest(h.logger, w, "failed to read the request payload", err)
		return
	}

	payload.CardID = cardID
	payload.ListID = listID

	if err := h.validator.Struct(payload); err != nil {
		helper.BadRequest(h.logger, w, "failed validation on the request payload", err)
		return
	}

	h.logger.Info("payload", zap.Any("payload", payload))

	if err := h.store.UpdateCard(r.Context(), cardID, &payload); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "card updated successfully", nil)
}

func (h *handler) handleGetCardDetail(w http.ResponseWriter, r *http.Request) {
	cardID := r.PathValue("cardID")
	card, err := h.store.GetCardDetail(r.Context(), cardID)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "card details fetched successfully", card)
}

func (h *handler) handleDeleteCard(w http.ResponseWriter, r *http.Request) {
	cardID := r.PathValue("cardID")
	if err := h.store.DeleteCard(r.Context(), cardID); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "card deleted successfully", nil)
}
