package libauth

import (
	"context"
	"net/http"
)

type SessionMiddleware struct {
	sessionStore SessionStore
}

func NewSessionMiddleware(sessionStore SessionStore) *SessionMiddleware {
	return &SessionMiddleware{sessionStore: sessionStore}
}

func (m *SessionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Session cookie missing", http.StatusUnauthorized)
			return
		}

		userID, err := m.sessionStore.Get(sessionID.Value)
		if err != nil {
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
