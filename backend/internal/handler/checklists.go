package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (h *handler) handleAddChecklistToCard(w http.ResponseWriter, r *http.Request) {
	cardID := r.PathValue("cardID")
	var createChecklist types.AddChecklist
	if err := helper.ReadJSON(r, &createChecklist); err != nil {
		helper.BadRequest(h.logger, w, "Invalid request payload", err)
		return
	}

	createChecklist.CardID = cardID

	if err := h.validator.Struct(createChecklist); err != nil {
		helper.BadRequest(h.logger, w, "Invalid request payload", err)
		return
	}

	if err := h.store.AddChecklistToCard(r.Context(), &createChecklist); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "Checklist added to card successfully", nil)
}

func (h *handler) handleGetChecklist(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("checklistID")
	if err := h.validator.Var(checklistID, "required,uuid"); err != nil {
		helper.BadRequest(h.logger, w, "invalid checklist id", nil)
		return
	}

	checklist, err := h.store.GetChecklist(r.Context(), checklistID)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "Checklist retrieved successfully", checklist)
}
