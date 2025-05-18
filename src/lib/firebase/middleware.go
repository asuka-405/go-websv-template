package firebase

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

type ContextKey string

const firebaseUserKey ContextKey = "firebaseUser"

func AuthMiddlware(authClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			if idToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			token, err := authClient.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), firebaseUserKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
