package middleware

import (
	"easyvpn/src/utils"
	"net/http"
)

func DecodeJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.HandleError("Token Was Not Provided", "DecodeJWT")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Verify the token
		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.HandleError("Token Is Invalid", "DecodeJWT")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
