package middleware

import (
	"easyvpn/src/utils"
	"fmt"
	"net/http"
	"strings"
)

func CheckUserRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.HandleError(fmt.Errorf("Token Was Not Provided"), "DecodeJWT")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := utils.VerifyToken(strings.Split(tokenString, "Bearer ")[1])
		if !token.TokenValid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CheckAdminRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.HandleError(fmt.Errorf("Token Was Not Provided"), "DecodeJWT")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := utils.VerifyToken(strings.Split(tokenString, "Bearer ")[1])
		if !token.TokenValid || !token.IsAdmin {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
