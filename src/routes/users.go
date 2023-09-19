package routes

import (
	"easyvpn/src/services"
	"encoding/json"
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

	return
}

func GetUsers(w http.ResponseWriter) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req dtos.UserID
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.DeleteUsers(req.ID)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dtos.FrontEndUsers
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var users []dtos.FrontEndUser
	for _, user := range req.Users {
		u, err := services.UpdateUser(user)
		if err != nil {
			utils.HandleError(err, "UpdateUser")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		users = append(users, services.FormatUser(*u))
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

func SetTemporaryPassword(w http.ResponseWriter, r *http.Request) {
	var req dtos.UserID
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "SetTemporaryPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.SetTempUserPassword(req.ID)
	if err != nil {
		utils.HandleError(err, "SetTemporaryPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

func ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	var req dtos.ChangePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "ChangeUserPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.ChangeUserPassword(req.ID, req.Password)
	if err != nil {
		utils.HandleError(err, "ChangeUserPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
