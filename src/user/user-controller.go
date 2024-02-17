package user

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel  `bun:"table:users,alias:u"`
	ID             uint      `bun:",pk,autoincrement" json:"id"`
	Name           string    `bun:",notnull" json:"name"`
	Username       string    `bun:",notnull" json:"username"`
	Password       string    `bun:",notnull" json:"password"`
	Roles          string    `bun:",notnull" json:"roles"`
	IsAdmin        bool      `json:"is_admin"`
	Enabled        bool      `json:"enabled"`
	PasswordExpiry time.Time `json:"password_expiry"`
}

const CompleteRoles = "Server Status,User Management,Settings"

func UsersPage(w http.ResponseWriter, r *http.Request) {
	username := "hello"
	users, err := GetUsers(r.URL.Query().Get("username"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Users(username, users, r.URL.Query().Get("username"), CompleteRoles).Render(r.Context(), w)
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
	user := User{
		Username:       r.Form.Get("username"),
		Password:       r.Form.Get("password"),
		IsAdmin:        r.Form.Get("admin") == "on",
		Enabled:        r.Form.Get("enabled") == "on",
		PasswordExpiry: passwordExpiry,
		Roles:          strings.Join(r.Form["roles"], ","),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)
	_, err = database.DB.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = utils.GenerateSignedCertificate(`C:\Program Files\OpenVPN\config-auto\keys\`, user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := database.DB.NewDelete().Model((*User)(nil)).Where("id = ?", chi.URLParam(r, "id")).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	users, err := GetUsers("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	UsersTable(users, r.URL.Query().Get("username"), CompleteRoles).Render(r.Context(), w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.Form.Encode())

	var passwordExpiry time.Time
	if r.Form.Get("mustChangePw") == "on" {
		passwordExpiry = time.Now()
	} else {
		passwordExpiry = time.Now().Add(30 * 24 * time.Hour)
	}
	user := User{
		Username:       r.Form.Get("username"),
		Password:       r.Form.Get("password"),
		IsAdmin:        r.Form.Get("admin") == "on",
		Enabled:        r.Form.Get("enabled") == "on",
		PasswordExpiry: passwordExpiry,
		Roles:          strings.Join(r.Form["roles"], ","),
	}
	_, err = database.DB.NewUpdate().Model(&user).Where("id = ?", chi.URLParam(r, "id")).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func GetUser(username string) (*User, error) {
	user := new(User)
	err := database.DB.NewSelect().Model(user).Where("username = ?", username).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers(username string) (*[]User, error) {
	users := new([]User)
	err := database.DB.NewSelect().Model(users).Scan(context.Background())
	if err != nil {
		return nil, err
	}

	var filteredusers []User
	for _, user := range *users {
		if fuzzy.Match(username, user.Username) {
			filteredusers = append(filteredusers, user)
		}
	}

	return &filteredusers, nil
}
