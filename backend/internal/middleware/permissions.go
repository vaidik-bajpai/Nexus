package middleware

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (m *Middleware) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := helper.GetUserFromRequestContext(r)
		boardID := r.PathValue("boardID")
		if err := m.validator.Var(boardID, "uuid"); err != nil || boardID == "" {
			helper.BadRequest(m.logger, w, "invalid boardID", nil)
			return
		}

		boardMember, err := m.store.GetBoardMember(r.Context(), boardID, user.ID)
		if err != nil {
			helper.InternalServerError(m.logger, w, nil, err)
			return
		}

		if boardMember.Role != "admin" {
			helper.Forbidden(m.logger, w, "you are forbidden to make this action", nil)
			return
		}

		next.ServeHTTP(w, helper.SetBoardInRequestContext(r, &types.Board{Name: boardMember.BoardName}))
	})
}
