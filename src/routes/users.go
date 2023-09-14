package routes

import (
	"easyvpn/src/services"
	"encoding/json"
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

	response, err := services.VerifyUser(requestData.Username, requestData.Password)
	if err != nil {
		utils.HandleError(err, "UserLogin")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}

func CheckUserToken(w http.ResponseWriter, r *http.Request) {
	var requestData dtos.CheckTokenRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.HandleError(err, "CheckUserToken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	check := utils.VerifyToken(requestData.Token)
	responseData := map[string]interface{}{}
	responseData["is_admin"] = check.IsAdmin
	responseData["token_valid"] = check.TokenValid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "CheckUserToken")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "CreateUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(req.Username, req.Name, req.Password, req.IsAdmin, req.Enabled)
	if err != nil {
		fmt.Println(err.Error())
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

	u := services.FormatUser(user)

	responseData := map[string]interface{}{}
	responseData["user"] = u

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "CreateUser")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		utils.HandleError(err, "GetUsers")
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
		utils.HandleError(err, "GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
