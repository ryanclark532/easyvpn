package routes

import (
	"database/sql"
	"easyvpn/src/services"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net/http"

	"easyvpn/src/dtos"
	"easyvpn/src/utils"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var requestData dtos.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.VerifyUser(requestData.Username, requestData.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		utils.HandleError(err, "UserLogin")
		return
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["token"] = token
	responseData["is_admin"] = true

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		return
	}

}

func CheckUserToken(w http.ResponseWriter, r *http.Request) {
	var requestData dtos.CheckTokenRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	check, err := utils.VerifyToken(requestData.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["is_admin"] = check.IsAdmin
	responseData["token_valid"] = check.TokenValid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(req.Username, req.Name, req.Password, req.IsAdmin, req.Enabled)
	if err != nil {
		fmt.Println(err.Error())
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				w.WriteHeader(http.StatusConflict)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	u := services.FormatUser(user)

	responseData := map[string]interface{}{}
	responseData["user"] = u

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u := services.FormatUsers(users)

	responseData := map[string]interface{}{}
	responseData["users"] = u
	responseData["count"] = len(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
