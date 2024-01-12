package settings

import (
	"easyvpn/src/logging"
	"easyvpn/src/settings/settings_dtos"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func GetSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	s, err := GetSettings()
	if err != nil {
		logging.HandleError(err, "GetSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		logging.HandleError(err, "GetSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SetSettingsEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *settings_dtos.Settings
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logging.HandleError(err, "PostSettingsEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(req)

	if err != nil {
		logging.HandleError(err, "PostSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetConfigFileEndpoint(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`)
	if err != nil {
		logging.HandleError(err, "GetConfigFile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(string(content))
	if err != nil {
		logging.HandleError(err, "GetSettings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ServerSettingsPage(w http.ResponseWriter, r *http.Request) {
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(settings)
	ServerSettings("test", settings).Render(r.Context(), w)
}

func ClientSettingsPage(w http.ResponseWriter, r *http.Request) {
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	ClientSettings("test", settings).Render(r.Context(), w)
}

func SetServerSettings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

	}
	Port, _ := strconv.ParseInt(r.Form.Get("vpn_port"), 10, 0)
	settings.Port = int(Port)
	WebServerPort, _ := strconv.ParseInt(r.Form.Get("web_port"), 10, 0)
	settings.WebServerPort = int(WebServerPort)
	MaxConnections, _ := strconv.ParseInt(r.Form.Get("max_connections"), 10, 0)
	settings.MaxConnections = int(MaxConnections)
	settings.UseAsGateway = r.Form.Get("use_as_gateway") == "on"

	err = SetSettings(settings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/settings/server", http.StatusSeeOther)
}
func SetClientSettings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(r.Form.Encode())
	settings.DNSServer1 = r.Form.Get("dns1")
	settings.DNSServer2 = r.Form.Get("dns2")
	settings.PrivateAccess = r.Form.Get("private_access") == "on"
	err = SetSettings(settings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/settings/client", http.StatusSeeOther)
}
