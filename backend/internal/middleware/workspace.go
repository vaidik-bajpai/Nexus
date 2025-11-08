package middleware

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (m *Middleware) VerifyWorkspaceAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		workspaceID := r.PathValue("workspace_id")
		if workspaceID == "" {
			helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
				Status:  http.StatusBadRequest,
				Message: "workspace id is required",
			})
		}

		user := r.Context().Value(types.UserCtxKey).(*types.User)
		if user.ID == "" {
			helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

	})
}
