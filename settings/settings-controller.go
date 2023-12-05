package settings

import (
	"easyvpn/settings/settings_dtos"
	"easyvpn/utils"
	"encoding/json"
	"net/http"
)

func GetSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	s, err := GetSettings()
	if err != nil {
		utils.HandleError(err, "GetSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		utils.HandleError(err, "GetSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}



func SetSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
var req *settings_dtos.Settings
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "PostSettingsEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SetSettings(req)
	if err != nil {
		utils.HandleError(err, "PostSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
	