package user

import (
	user_dtos "easyvpn/src/user/user-dtos"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pkgz/auth/token"
	"github.com/go-sql-driver/mysql"

	"easyvpn/src/utils"
)

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *user_dtos.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "CreateUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := CreateUser(req)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				utils.HandleError(err, "CreateUser")
				w.WriteHeader(http.StatusConflict)
				return
			}
			utils.HandleError(err, "CreateUser")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			utils.HandleError(err, "CreateUser")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	frontendUser := FormatUser(user)

	responseData := map[string]interface{}{}
	responseData["user"] = frontendUser

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "CreateUser")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}

func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	xu, err := token.GetUserInfo(r)
	fmt.Println(err)
	fmt.Println(xu)
	users, err := GetUsers()
	if err != nil {
		utils.HandleError(err, "GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u := FormatUsers(*users)

	responseData := map[string]interface{}{}
	responseData["users"] = u
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *user_dtos.UserID
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DeleteUsers(req.ID)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *user_dtos.FrontEndUsers
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var users []user_dtos.FrontEndUser
	for _, user := range req.Users {
		u, err := UpdateUser(user)
		if err != nil {
			utils.HandleError(err, "UpdateUser")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		users = append(users, FormatUser(u))
	}

	responseData := map[string]interface{}{}
	responseData["users"] = users

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "UpdateUser")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
