package user

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/logging"
	user_dtos "easyvpn/src/user/user-dtos"
	"easyvpn/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func SetPWEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *user_dtos.PasswordResetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SetPassword(chi.URLParam(r, "id"), req)
	if err != nil {
		logging.HandleError(err, "SetPWEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CreateUserConfigEndpoint(w http.ResponseWriter, r *http.Request) {
	err := utils.GenerateClientConfig(chi.URLParam(r, "username"))
	if err != nil {
		logging.HandleError(err, "CreateUserConfig")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUserConfigEndpoint(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(fmt.Sprintf(`./tmp/%s.ovpn`, chi.URLParam(r, "username")))
	if err != nil {
		logging.HandleError(err, "DeleteUserConfig")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UsersPage(w http.ResponseWriter, r *http.Request) {
	username := "hello"
	users, err := GetUsers(r.URL.Query().Get("username"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Users(username, users).Render(r.Context(), w)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	var passwordExpiry time.Time
	if r.Form.Get("mustChangePw") == "on" {
		passwordExpiry = time.Now()
	} else {
		passwordExpiry = time.Now().Add(30 * 24 * time.Hour)
	}
	user := user_dtos.User{
		Username:       r.Form.Get("username"),
		Password:       r.Form.Get("password"),
		IsAdmin:        r.Form.Get("admin") == "on",
		Enabled:        r.Form.Get("enabled") == "on",
		PasswordExpiry: passwordExpiry,
	}

	err = CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := database.DB.NewDelete().Model((*user_dtos.User)(nil)).Where("id = ?", chi.URLParam(r, "id")).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	users, err := GetUsers("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	UsersTable(users).Render(r.Context(), w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	var passwordExpiry time.Time
	if r.Form.Get("mustChangePw") == "on" {
		passwordExpiry = time.Now()
	} else {
		passwordExpiry = time.Now().Add(30 * 24 * time.Hour)
	}
	user := user_dtos.User{
		Username:       r.Form.Get("username"),
		Password:       r.Form.Get("password"),
		IsAdmin:        r.Form.Get("admin") == "on",
		Enabled:        r.Form.Get("enabled") == "on",
		PasswordExpiry: passwordExpiry,
	}
	_, err = database.DB.NewUpdate().Model(user).Where("id = ?", chi.URLParam(r, "id")).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users, err := GetUsers("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	UsersTable(users).Render(r.Context(), w)
}
