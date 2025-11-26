package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (h *handler) handleInviteToBoard(w http.ResponseWriter, r *http.Request) {
	user := helper.GetUserFromRequestContext(r)
	board := helper.GetBoardFromRequestContext(r)
	boardID := r.PathValue("boardID")

	var payload *types.BoardInvitation
	if err := helper.ReadJSON(r, &payload); err != nil {
		helper.UnprocessableEntity(h.logger, w, "invalid request payload", nil)
		return
	}

	payload.BoardID = boardID
	payload.InvitedBy = user.ID
	payload.Token = strconv.Itoa(helper.CreateRandomToken())
	payload.ExpiredAt = time.Now().Add(2 * 24 * time.Hour)

	if err := h.validator.Struct(payload); err != nil {
		helper.BadRequest(h.logger, w, "validation on request payload", nil)
		return
	}

	if isMember, err := h.store.IsABoardMember(r.Context(), payload.Email, payload.BoardID); !isMember {
		if err != nil {
			helper.InternalServerError(h.logger, w, nil, err)
			return
		}
		helper.Conflict(h.logger, w, "member with this email already exists", nil)
		return
	}

	err := h.store.CreateBoardInvitation(r.Context(), payload)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	if err = h.mailer.SendBoardInvitationEmail(
		[]string{payload.Email},
		"Invitation to join nexus",
		user.Username,
		board.Name,
		fmt.Sprintf("http://localhost:3000/join?token=%s", payload.Token),
	); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "invitation email sent successfully", nil)
}

func (h *handler) handleAcceptInviteToBoard(w http.ResponseWriter, r *http.Request) {
	user := helper.GetUserFromRequestContext(r)
	token := r.URL.Query().Get("token")
	if token == "" {
		helper.BadRequest(h.logger, w, "invalid request token", nil)
		return
	}

	invitation, err := h.store.GetBoardInvitationByToken(r.Context(), token)
	if err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	if invitation.ExpiredAt.Before(time.Now()) {
		helper.BadRequest(h.logger, w, "invitation expired", nil)
		return
	}

	if invitation.Email != user.Email {
		helper.BadRequest(h.logger, w, "invalid request", nil)
		return
	}

	if err := h.store.AcceptBoardInvitation(r.Context(), invitation.Token, user.ID, invitation.Role); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "invitation accepted successfully", nil)
}

func (h *handler) handleUpdateBoard(w http.ResponseWriter, r *http.Request) {
	boardID := r.PathValue("boardID")
	var payload *types.UpdateBoard
	if err := helper.ReadJSON(r, &payload); err != nil {
		helper.UnprocessableEntity(h.logger, w, "invalid request payload", nil)
		return
	}

	if err := h.validator.Struct(payload); err != nil {
		helper.BadRequest(h.logger, w, "validation on request payload", nil)
		return
	}

	payload.BoardID = boardID

	if err := h.store.UpdateBoard(r.Context(), payload); err != nil {
		helper.InternalServerError(h.logger, w, nil, err)
		return
	}

	helper.OK(h.logger, w, "board updated successfully", nil)
}
