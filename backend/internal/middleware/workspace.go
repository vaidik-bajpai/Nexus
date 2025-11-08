package middleware

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (m *Middleware) VerifyWorkspaceAccess(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			workspaceID := r.PathValue("workspace_id")
			if workspaceID == "" {
				helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
					Status:  http.StatusBadRequest,
					Message: "workspace id is required",
				})
				return
			}

			user, ok := r.Context().Value(types.UserCtxKey).(*types.User)
			if !ok || user == nil || user.ID == "" {
				helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
				return
			}

			workspace, err := m.store.GetWorkspace(r.Context(), workspaceID)
			if err != nil {
				helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
					Status:  http.StatusInternalServerError,
					Message: "failed to get workspace",
				})
				return
			}

			for _, member := range workspace.Members {
				if member.UserID == user.ID {
					memberRoleLevel, ok := types.WorkspaceRoles[member.Role]
					if !ok {
						continue
					}
					requiredRoleLevel, ok := types.WorkspaceRoles[requiredRole]
					if !ok {
						helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
							Status:  http.StatusInternalServerError,
							Message: "invalid role specified",
						})
						return
					}
					// Lower number = higher privilege (admin=1, manager=2, member=3)
					if memberRoleLevel <= requiredRoleLevel {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			helper.WriteJSON(w, http.StatusForbidden, &types.Response{
				Status:  http.StatusForbidden,
				Message: "insufficient permissions",
			})
		})
	}
}

func (m *Middleware) RequireAdmin() func(http.Handler) http.Handler {
	return m.VerifyWorkspaceAccess("admin")
}

func (m *Middleware) RequireManager() func(http.Handler) http.Handler {
	return m.VerifyWorkspaceAccess("manager")
}

func (m *Middleware) RequireMember() func(http.Handler) http.Handler {
	return m.VerifyWorkspaceAccess("member")
}
