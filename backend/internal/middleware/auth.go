package middleware

import (
	"context"
	"net/http"
	"strings"

	"log"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func (m *Middleware) VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("verifying access token")

		token := r.Header.Get("Authorization")
		if token == "" {
			helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		if token == "" {
			helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

		user, err := helper.VerifyAccessToken(token)
		if err != nil {
			log.Println("error verifying access token", err)
			helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

		usr, err := m.store.GetUserByEmail(r.Context(), user.Email)
		if err != nil {
			log.Println("error getting user by email", err)
			helper.WriteJSON(w, http.StatusUnauthorized, &types.Response{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

		log.Println("user", usr)

		ctx := context.WithValue(r.Context(), types.UserCtxKey, usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
