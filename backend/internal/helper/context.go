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

func SetPaginateInRequestContext(r *http.Request, paginate *types.Paginate) *http.Request {
	ctx := context.WithValue(r.Context(), types.PaginateCtxKey, paginate)
	return r.WithContext(ctx)
}

func GetPaginateFromRequestContext(r *http.Request) *types.Paginate {
	paginate, ok := r.Context().Value(types.PaginateCtxKey).(*types.Paginate)
	if !ok || paginate == nil {
		return types.DefaultPaginate()
	}

	return paginate
}
