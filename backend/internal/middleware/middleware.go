package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
	"go.uber.org/zap"
)

type Middleware struct {
	store     store.Storer
	validator *validator.Validate
	logger    *zap.Logger
}

func NewMiddleware(store *store.Store, l *zap.Logger, v *validator.Validate) *Middleware {
	return &Middleware{store: store, validator: v, logger: l}
}
