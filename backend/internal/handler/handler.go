package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	m "github.com/vaidik-bajpai/Nexus/backend/internal/middleware"
)

type handler struct {
	logger     *zap.Logger
	validator  *validator.Validate
	store      *store.Store
	oauth2     map[string]*oauth2.Config
	middleware *m.Middleware
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
		logger:     l,
		validator:  v,
		store:      store,
		oauth2:     oauth2Configs,
		middleware: m.NewMiddleware(store),
	}
}

func (h *handler) SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/register", h.handleUserRegistration)
		r.Post("/login", h.handleUserLogin)
		r.Get("/{provider}", h.handleUserOAuthFlow)
		r.Get("/{provider}/callback", h.handleUserOAuthCallback)
		r.With(h.middleware.VerifyAccessToken).Post("/logout", h.handleUserLogout)
		r.Post("/reset-password", h.handlePasswordResetFlow)
		r.Post("/password/reset", h.handlePasswordReset)
		r.Post("/refresh-token", h.handleRefreshToken)
	})

	r.Route("/api/v1/workspace", func(r chi.Router) {
		r.With(h.middleware.VerifyAccessToken).Post("/create", h.handleCreateWorkspace)
		r.With(h.middleware.VerifyAccessToken).Get("/list", h.handleListWorkspaces)

		r.Route("/{workspace_id}", func(r chi.Router) {
			r.Use(h.middleware.VerifyAccessToken)
			r.Get("/", h.handleGetWorkspace)
			r.With(h.middleware.RequireMember()).Get("/project/list", h.handleListProjects)
			r.With(h.middleware.RequireManager()).Post("/project/create", h.handleCreateProject)
		})
	})

	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Use(h.middleware.VerifyAccessToken)
		r.Route("/{project_id}", func(r chi.Router) {
			r.With(h.middleware.VerifyProjectAccess("manager")).Post("/create", h.handleCreateTask)
			r.With(h.middleware.VerifyProjectAccess("member")).Get("/list", h.handleListTasks)
		})
	})

	return r
}
