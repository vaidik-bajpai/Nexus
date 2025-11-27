package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
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
