package libauth

import (
	"context"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	authService *AuthService
}

func NewJWTMiddleware(authService *AuthService) *JWTMiddleware {
	return &JWTMiddleware{authService: authService}
}

func (m *JWTMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := m.authService.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
