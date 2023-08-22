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

	responseData := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError("Error Encoding Response", "UserLogin")
		return
	}

}
