package middleware

import (
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

// Paginate is a middleware that parses pagination query parameters
// and populates the request context with pagination information.
// Query parameters:
//   - page: page number (default: 1)
//   - size: page size (default: 10, max: 100)
//   - sort_by: field to sort by (default: "created_at")
//   - sort_order: "asc" or "desc" (default: "desc")
//
// The pagination object is available in handlers via:
//
//	paginate := helper.GetPaginateFromRequestContext(r)
func (m *Middleware) Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		pageStr := query.Get("page")
		sizeStr := query.Get("size")
		sortBy := query.Get("sort_by")
		sortOrder := query.Get("sort_order")

		paginate := types.ParsePaginateFromQuery(pageStr, sizeStr, sortBy, sortOrder)

		next.ServeHTTP(w, helper.SetPaginateInRequestContext(r, paginate))
	})
}
