package middleware

import (
	"easyvpn/src/utils"
	"fmt"
	"net/http"
)

func DecodeJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.HandleError(fmt.Errorf("Token Was Not Provided"), "DecodeJWT")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.HandleError(fmt.Errorf("Token Is Invalid"), "DecodeJWT")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
