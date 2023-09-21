package auth

import (
	"easyvpn/src/auth/auth_dtos"
	"easyvpn/src/utils"
	"encoding/json"
	"net/http"
)

func UserLoginEndpoint(w http.ResponseWriter, r *http.Request) {
	var requestData *auth_dtos.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := UserLogin(requestData)
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

func CheckUserTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var requestData *auth_dtos.CheckTokenRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.HandleError(err, "CheckUserToken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	check, err := utils.CheckUserToken(requestData.Token)
	if err != nil {
		utils.HandleError(err, "CheckUserToken")
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
		utils.HandleError(err, "CheckUserToken")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SetTemporaryPasswordEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *auth_dtos.UserID
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "SetTemporaryPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SetTempUserPassword(req.ID)
	if err != nil {
		utils.HandleError(err, "SetTemporaryPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

func ChangeUserPasswordEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *auth_dtos.ChangePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "ChangeUserPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ChangeUserPassword(req)
	if err != nil {
		utils.HandleError(err, "ChangeUserPassword")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
