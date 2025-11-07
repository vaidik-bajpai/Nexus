package middleware

import (
	"context"
	"net/http"
	"strings"

	"log"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func VerifyAccessToken(next http.Handler) http.Handler {
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

		log.Println("user", user)

		ctx := context.WithValue(r.Context(), types.UserCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
