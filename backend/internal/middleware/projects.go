package middleware

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (m *Middleware) VerifyProjectAccess(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(types.UserCtxKey).(*types.User)
			if !ok || user == nil || user.ID == "" {
				helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
				return
			}

			projectID := r.PathValue("project_id")
			if projectID == "" {
				helper.WriteJSON(w, http.StatusBadRequest, &types.Response{
					Status:  http.StatusBadRequest,
					Message: "project id is required",
				})
				return
			}

			project, err := m.store.GetProject(r.Context(), projectID)
			if err != nil {
				helper.WriteJSON(w, http.StatusInternalServerError, &types.Response{
					Status:  http.StatusInternalServerError,
					Message: "failed to get project",
				})
				return
			}

			workspace, err := m.store.GetWorkspace(r.Context(), project.WorkspaceID)
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
