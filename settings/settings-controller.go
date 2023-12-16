package settings

import (
	"easyvpn/settings/settings_dtos"
	"easyvpn/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(utils.Settings)
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
	fmt.Println(req)

	err = SetSettings(req)
	if err != nil {
		utils.HandleError(err, "PostSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetConfigFileEndpoint(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`)
	if err != nil {
		utils.HandleError(err, "GetConfigFile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(string(content))
	if err != nil {
		utils.HandleError(err, "GetSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
