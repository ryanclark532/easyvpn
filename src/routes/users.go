package routes

import (
	"encoding/json"
	"net/http"

	"easyvpn/src/dtos"
	"easyvpn/src/utils"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var requestData dtos.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.HandleError("Bad Request", "UserLogin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Handle Auth stuff here

	userID := int64(123)
	token, err := utils.CreateToken(userID)
	if err != nil {
		utils.HandleError("Error Creating Token", "UserLogin")
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
		utils.HandleError(err.Error(), "UserLogin")
		return
	}

}

func CheckUserToken(w http.ResponseWriter, r *http.Request) {
	var requestData dtos.CheckTokenRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.HandleError(err.Error(), "CheckUserToken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	check, err := utils.VerifyToken(requestData.Token)
	if err != nil {
		utils.HandleError(err.Error(), "CheckUserToken")
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
		utils.HandleError(err.Error(), "CheckUserToken")
		return
	}
}
