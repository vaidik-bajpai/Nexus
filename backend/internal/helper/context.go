package helper

import (
	"context"
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func SetUserInRequestContext(r *http.Request, user *types.User) *http.Request {
	ctx := context.WithValue(r.Context(), types.UserCtxKey, user)
	return r.WithContext(ctx)
}

func GetUserFromRequestContext(r *http.Request) *types.User {
	user, ok := r.Context().Value(types.UserCtxKey).(*types.User)
	if !ok || user == nil || user.ID == "" {
		panic("user not found in context")
	}

	return user
}

func SetProjectInRequestContext(r *http.Request, project *types.Project) *http.Request {
	ctx := context.WithValue(r.Context(), types.ProjectCtxKey, project)
	return r.WithContext(ctx)
}

func GetProjectFromRequestContext(r *http.Request) *types.Project {
	project, ok := r.Context().Value(types.ProjectCtxKey).(*types.Project)
	if !ok || project == nil || project.ID == "" {
		panic("project not found in context")
	}

	return project
}
