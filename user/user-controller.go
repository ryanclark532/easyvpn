package user

import (
	user_dtos "easyvpn/user/user-dtos"
	"easyvpn/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
)

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *user_dtos.User
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

	frontendUser := FormatUser(*user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(frontendUser)
	if err != nil {
		utils.HandleError(err, "CreateUser")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {

	users, err := GetUsers()
	if err != nil {
		utils.HandleError(err, "GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u := FormatUsers(*users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		utils.HandleError(err, "GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	err := DeleteUser(chi.URLParam(r, "id"))
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *user_dtos.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	 err = UpdateUser(req, chi.URLParam(r, "id"))
	if err !=nil {
		utils.HandleError(err, "UpdateUser")
	}
	w.WriteHeader(http.StatusNoContent)
}


func SetPWEndpoint(w http.ResponseWriter, r *http.Request){
	var req *user_dtos.PasswordResetRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SetPassword(chi.URLParam(r, "id"), req)
	if err !=nil {
		utils.HandleError(err, "SetPWEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
