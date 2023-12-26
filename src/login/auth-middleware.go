package login

import (
	"context"
	"easyvpn/src/user"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString, err := r.Cookie("JWT")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			}
			return []byte("YOUR_SECRET_KEY"), nil
		})

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				username, ok := claims["username"].(string)
				if !ok {
					http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
					return
				}
				user, err := user.GetUser(username)
				if err != nil {
					http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
					return
				}
				ctx := context.WithValue(r.Context(), "user", user)

				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
				return
			} else {
				http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
				return
			}
		} else {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
	})
}

func decodeJWT(tokenString string) (*jwt.Token, error) {
	secret := []byte("your-secret-key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	return token, err
}
