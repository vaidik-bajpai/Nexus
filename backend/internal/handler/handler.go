package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
	"go.uber.org/zap"
)

type handler struct {
	logger    *zap.Logger
	validator *validator.Validate
	store     *store.Store
}

func NewHandler(store *store.Store) *handler {
	l, _ := zap.NewDevelopment()
	v := validator.New()

	return &handler{
		logger:    l,
		validator: v,
		store:     store,
	}
}

func (h *handler) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users/register", h.handleUserRegistration)
	mux.HandleFunc("/api/v1/users/login", h.handleUserLogin)

	return mux
}
