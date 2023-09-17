package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allow specified HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
