package middleware

import (
	"context"
	"net/http"
	"strings"

	"log/slog"

	jwtService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/jwt"
)

func JWTAuth(log *slog.Logger, jwtService *jwtService.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			claims, err := jwtService.ParseToken(parts[1])
			if err != nil {
				log.Error("invalid token", "err", err)
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
