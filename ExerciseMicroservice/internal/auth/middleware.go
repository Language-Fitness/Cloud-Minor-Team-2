package auth

import (
	"context"
	"net/http"
	"strings"
)

var TokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

// Middleware extracts the Authorization Bearer token from the request headers and packs it into the context.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the Authorization header
		authHeader := r.Header.Get("Authorization")

		// Extract the Bearer token from the Authorization header
		var token string
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}

		// Store the token in the context
		ctx := context.WithValue(r.Context(), TokenCtxKey, token)

		// Call the next with our new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TokenFromContext retrieves the Bearer token from the context.
// REQUIRES Middleware to have run.
func TokenFromContext(ctx context.Context) string {
	raw, _ := ctx.Value(TokenCtxKey).(string)
	return raw
}
