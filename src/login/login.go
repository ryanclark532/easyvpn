package login

import (
	"easyvpn/src/user"
	"fmt"
	"net/http"

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

	ok, err := AuthUser(username, password)
	if err != nil {
		loginFormWithError(err.Error()).Render(r.Context(), w)
		return
	}
	if !ok {
		loginFormWithError("The Username or Password are incorrect").Render(r.Context(), w)
		return
	}
	loginForm().Render(r.Context(), w)
}

func AuthUser(username string, password string) (bool, error) {
	user, err := user.GetUser(username)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
