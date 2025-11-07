package middleware

import "github.com/vaidik-bajpai/Nexus/backend/internal/store"

type Middleware struct {
	store *store.Store
}

func NewMiddleware(store *store.Store) *Middleware {
	return &Middleware{store: store}
}
