package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	m "github.com/vaidik-bajpai/Nexus/backend/internal/middleware"
)

type handler struct {
	logger    *zap.Logger
	validator *validator.Validate
	store     *store.Store
	oauth2    map[string]*oauth2.Config
}

func NewHandler(store *store.Store) *handler {
	l, _ := zap.NewDevelopment()
	v := validator.New()

	oauth2Configs := make(map[string]*oauth2.Config)
	redirectURL := helper.GetStrEnvOrPanic("GOOGLE_REDIRECT_URL")

	oauth2Configs["google"] = &oauth2.Config{
		ClientID:     helper.GetStrEnvOrPanic("GOOGLE_CLIENT_ID"),
		ClientSecret: helper.GetStrEnvOrPanic("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return &handler{
		logger:    l,
		validator: v,
		store:     store,
		oauth2:    oauth2Configs,
	}
}

func (h *handler) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users/register", h.handleUserRegistration)
	mux.HandleFunc("POST /api/v1/users/login", h.handleUserLogin)
	mux.HandleFunc("GET /api/v1/users/{provider}", h.handleUserOAuthFlow)
	mux.HandleFunc("GET /api/v1/users/{provider}/callback", h.handleUserOAuthCallback)
	mux.Handle("POST /api/v1/users/logout", m.VerifyAccessToken(http.HandlerFunc(h.handleUserLogout)))

	return mux
}
