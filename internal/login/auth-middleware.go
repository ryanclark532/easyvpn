package login

import (
	"context"
	"easyvpn/internal/user"
	"fmt"
	"net/http"
	"strings"

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
				fmt.Println(user.Roles)

				if !checkRoles(r.URL.Path, user.Roles) {
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

func checkRoles(route string, roles string) bool {
	if strings.HasPrefix(route, "/vpn") {
		return strings.Contains(roles, "Server Status")
	}
	if strings.HasPrefix(route, "/settings") {
		return strings.Contains(roles, "Settings")
	}
	if strings.HasPrefix(route, "/users") || strings.HasPrefix(route, "/groups") {
		return strings.Contains(roles, "User Management")
	}
	return false
}
