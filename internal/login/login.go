package login

import (
	"easyvpn/internal/user"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		loginFormWithError(err.Error()).Render(r.Context(), w)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	ok, user, err := AuthUser(username, password)
	if err != nil {
		loginFormWithError(err.Error()).Render(r.Context(), w)
		return
	}
	if !ok {
		loginFormWithError("The Username or Password are incorrect").Render(r.Context(), w)
		return
	}

	jwt, err := generateJWT(user)
	if err != nil {
		loginFormWithError(err.Error()).Render(r.Context(), w)
		return
	}

	cookie := http.Cookie{
		Name:     "JWT",
		Value:    jwt,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func HandleSignout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "JWT",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AuthUser(username string, password string) (bool, *user.User, error) {
	user, err := user.GetUser(username)
	if err != nil {
		return false, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil, err
	}
	return true, user, nil
}

func generateJWT(user *user.User) (string, error) {
	secret := []byte("YOUR_SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["roles"] = user.Roles
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
