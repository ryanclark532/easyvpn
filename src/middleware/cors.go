package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupCORS(r *mux.Router) http.Handler {
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)
	return corsHandler
}
