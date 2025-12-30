package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/mailer"
	m "github.com/vaidik-bajpai/Nexus/backend/internal/middleware"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

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
			r.With(h.middleware.VerifyAccessToken).Get("/me", h.handleGetMe)
		})

		r.Route("/boards", func(r chi.Router) {
			r.Use(h.middleware.VerifyAccessToken)
			r.Post("/create", h.handleCreateBoard)

			r.With(h.middleware.Paginate).Get("/list", h.handleListBoards)

			r.Route("/{boardID}", func(r chi.Router) {
				r.Post("/accept-invite", h.handleAcceptInviteToBoard)

				r.Group(func(r chi.Router) {
					r.Use(h.middleware.IsMember)
					r.Get("/cards-and-lists", h.handleGetCardsAndLists)
					r.Get("/details", h.handleGetBoardDetails)
				})

				r.Group(func(r chi.Router) {
					r.Use(h.middleware.IsAdmin)
					r.Post("/invite", h.handleInviteToBoard)
					r.Put("/update", h.handleUpdateBoard)
					r.Delete("/delete", h.handleDeleteBoard)
				})

				r.Route("/labels", func(r chi.Router) {
					r.Use(h.middleware.IsMember)
					r.Post("/create", h.handleCreateLabel)
					r.Post("/modify", h.handleModifyLabel)
					r.Get("/list", h.handleListBoardLabels)
				})

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
								r.Get("/detail", h.handleGetCardDetail)
								r.Delete("/delete", h.handleDeleteCard)
								r.Post("/toggle-member", h.handleToggleCardMembership)

								r.Route("/labels", func(r chi.Router) {
									r.Post("/toggle", h.handleToggleLabelToCard)
									r.Post("/list", h.handleListCardLabels)
								})

								r.Route("/checklists", func(r chi.Router) {
									r.Post("/create", h.handleAddChecklistToCard)
									r.Route("/{checklistID}", func(r chi.Router) {
										r.Get("/detail", h.handleGetChecklist)
										r.Delete("/delete", h.handleDeleteChecklist)

										r.Route("/items", func(r chi.Router) {
											r.Post("/create", h.handleAddChecklistItem)
											r.Route("/{itemID}", func(r chi.Router) {
												r.Delete("/delete", h.handleDeleteChecklistItem)
												r.Put("/update", h.handleUpdateChecklistItem)
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})

	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "uploads"))
	FileServer(r, "/uploads", filesDir)

	r.Post("/api/v1/upload", h.handleUpload)

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
