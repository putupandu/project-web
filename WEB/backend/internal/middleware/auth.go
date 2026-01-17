package middleware

import (
	"context"
	"e-library/backend/internal/models"
	"net/http"
)
//
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Untuk development: gunakan user dummy
		dummyUser := &models.User{
			ID:    1,
			Email: "dummy@example.com",
			Name:  "Dummy User",
		}
		ctx := context.WithValue(r.Context(), "user", dummyUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}