package vpn

import (
	"easyvpn/src/utils"
	"encoding/json"
	"net/http"
)

func GetServerStatusEndpoint(w http.ResponseWriter, r *http.Request) {
	status, err := GetVpnServerStatus()
	if err != nil {
		utils.HandleError(err, "GetServerStatusEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["status"] = status

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "GetServerStatusEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
