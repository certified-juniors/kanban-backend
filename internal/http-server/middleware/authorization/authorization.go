package authorization

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	customjwt "kanban/internal/lib/jwt"
	"net/http"
	"strings"
)

func AuthorizationMiddleware(secret string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			token, err := customjwt.ValidateToken(tokenString, []byte(secret))
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			customClaims, ok := claims["custom_claims"].(map[string]interface{})
			if !ok {
				http.Error(w, "Invalid custom claims", http.StatusUnauthorized)
				return
			}

			userIDStr, ok := customClaims["id"].(string)
			if !ok {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				http.Error(w, "Invalid UUID format", http.StatusUnauthorized)
				return
			}

			// TODO: Make all keys to constants (not hardcoded)
			ctx := context.WithValue(r.Context(), "id", userID)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
