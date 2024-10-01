package middleware

import (
	"github.com/ryanjarv/yxks/pkg/utils"
	"net/http"
	"strings"
)

// SuperSecureMiddleware makes everything secure
func SuperSecureMiddleware(ctx utils.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details

		userAgent := r.Header.Get("User-Agent")
		if !strings.HasPrefix(userAgent, "KMS-External-Key-Store/") {
			w.WriteHeader(http.StatusForbidden)
			ctx.Info.Printf("Access denied: %s %s from %s", r.Method, r.URL, r.RemoteAddr)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
