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

func (h *handler) handleDeleteChecklist(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("checklistID")
	if err := h.validator.Var(checklistID, "required,uuid"); err != nil {
		helper.BadRequest(h.logger, w, "invalid checklist id", nil)
		return
	}

	if err := h.store.DeleteChecklist(r.Context(), checklistID); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "Checklist deleted successfully", nil)
}

func (h *handler) handleAddChecklistItem(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("checklistID")
	var addItem types.AddChecklistItem
	if err := helper.ReadJSON(r, &addItem); err != nil {
		helper.BadRequest(h.logger, w, "Invalid request payload", err)
		return
	}

	addItem.ChecklistID = checklistID

	if err := h.validator.Struct(addItem); err != nil {
		helper.BadRequest(h.logger, w, "Invalid request payload", err)
		return
	}

	item, err := h.store.AddChecklistItem(r.Context(), &addItem)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "Checklist item added successfully", item)
}

func (h *handler) handleDeleteChecklistItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemID")
	if err := h.validator.Var(itemID, "required,uuid"); err != nil {
		helper.BadRequest(h.logger, w, "invalid item id", nil)
		return
	}

	if err := h.store.DeleteChecklistItem(r.Context(), itemID); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "Checklist item deleted successfully", nil)
}

func (h *handler) handleUpdateChecklistItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemID")
	var updateItem types.UpdateChecklistItem
	if err := helper.ReadJSON(r, &updateItem); err != nil {
		helper.BadRequest(h.logger, w, "Invalid request payload", err)
		return
	}

	updateItem.ItemID = itemID

	if err := h.validator.Struct(updateItem); err != nil {
		helper.BadRequest(h.logger, w, "Invalid request payload", err)
		return
	}

	if err := h.store.UpdateChecklistItem(r.Context(), &updateItem); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "Checklist item updated successfully", nil)
}
