package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (h *handler) handleCreateList(w http.ResponseWriter, r *http.Request) {
	var payload *types.CreateList
	if err := helper.ReadJSON(r, &payload); err != nil {
		helper.UnprocessableEntity(h.logger, w, "invalid request payload", nil)
		return
	}

	boardID := r.PathValue("boardID")
	payload.BoardID = boardID

	if err := h.validator.Struct(payload); err != nil {
		helper.BadRequest(h.logger, w, "failed validation on the request payload", nil)
		return
	}

	if err := h.store.CreateList(r.Context(), payload); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "list created successfully", nil)
}

func (h *handler) handleUpdateList(w http.ResponseWriter, r *http.Request) {
	var payload *types.UpdateLists
	if err := helper.ReadJSON(r, &payload); err != nil {
		helper.UnprocessableEntity(h.logger, w, "invalid request payload", nil)
		return
	}

	listID := r.PathValue("listID")
	payload.ListID = listID

	if err := h.validator.Struct(payload); err != nil {
		helper.BadRequest(h.logger, w, "failed validation on the request payload", nil)
		return
	}

	if err := h.store.UpdateList(r.Context(), payload); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "list updated successfully", nil)
}
