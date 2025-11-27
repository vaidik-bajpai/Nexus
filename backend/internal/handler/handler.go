package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/mailer"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	m "github.com/vaidik-bajpai/Nexus/backend/internal/middleware"
)

type handler struct {
	logger     *zap.Logger
	validator  *validator.Validate
	store      store.Storer
	mailer     mailer.Mailer
	oauth2     map[string]*oauth2.Config
	middleware *m.Middleware
}

func NewHandler(store *store.Store) *handler {
	l, _ := zap.NewDevelopment()
	v := validator.New()
	if err := helper.RegisterCustomValidations(v); err != nil {
		panic(err)
	}

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
		mailer:     mailer.NewSMTPMailer(),
		oauth2:     oauth2Configs,
		middleware: m.NewMiddleware(store, l, v),
	}
}

func (h *handler) SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/api/v1/", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", h.handleUserRegistration)
			r.Post("/login", h.handleUserLogin)
			r.Get("/{provider}", h.handleUserOAuthFlow)
			r.Get("/{provider}/callback", h.handleUserOAuthCallback)
			r.With(h.middleware.VerifyAccessToken).Post("/logout", h.handleUserLogout)
			r.Post("/reset-password", h.handlePasswordResetFlow)
			r.Post("/password/reset", h.handlePasswordReset)
			r.Post("/refresh-token", h.handleRefreshToken)
		})

		r.Route("/boards", func(r chi.Router) {
			r.Use(h.middleware.VerifyAccessToken)
			r.Post("/create", h.handleCreateBoard)

			r.With(h.middleware.Paginate).Get("/list", h.handleListBoards)

			r.Route("/{boardID}", func(r chi.Router) {
				r.With(h.middleware.IsAdmin).Post("/invite", h.handleInviteToBoard)
				r.Post("/accept-invite", h.handleAcceptInviteToBoard)
				r.With(h.middleware.IsAdmin).Put("/update", h.handleUpdateBoard)
				r.With(h.middleware.IsAdmin).Delete("/delete", h.handleDeleteBoard)

				r.Route("/lists", func(r chi.Router) {
					r.Use(h.middleware.IsMember)
					r.Post("/create", h.handleCreateList)
					r.Route("/{listID}", func(r chi.Router) {
						r.Put("/update", h.handleUpdateList)
						r.Delete("/delete", h.handleDeleteList)

						r.Route("/cards", func(r chi.Router) {
							r.Post("/create", h.handleCreateCard)
							r.Route("/{cardID}", func(r chi.Router) {
								r.Put("/update", h.handleUpdateCard)
							})
						})
					})
				})
			})
		})

	})

	return r
}
