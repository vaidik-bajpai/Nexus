package handler

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (h *handler) handleCreateBoard(w http.ResponseWriter, r *http.Request) {
	var board types.CreateBoard
	if err := helper.ReadJSON(r, &board); err != nil {
		helper.UnprocessableEntity(h.logger, w, "invalid request payload", nil)
		return
	}

	if err := h.validator.Struct(board); err != nil {
		helper.BadRequest(h.logger, w, "failed validation on the request payload", nil)
		return
	}

	owner := helper.GetUserFromRequestContext(r)
	board.OwnerID = owner.ID

	if err := h.store.CreateBoard(r.Context(), &board); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.Created(h.logger, w, "board created successfully", nil)
}

func (h *handler) handleListBoards(w http.ResponseWriter, r *http.Request) {
	user := helper.GetUserFromRequestContext(r)
	paginate := helper.GetPaginateFromRequestContext(r)

	boards, err := h.store.ListBoards(r.Context(), user.ID, paginate)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "boards fetched successfully", map[string]any{"boards": boards})
}
