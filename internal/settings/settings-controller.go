package settings

import (
	"easyvpn/internal/common"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

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

func AuthSettingsPage(w http.ResponseWriter, r *http.Request) {
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	AuthSettings("test", settings).Render(r.Context(), w)
}

func ConfigFileSettingsPage(w http.ResponseWriter, r *http.Request) {
	contents, err := os.ReadFile(common.VPN_TCP_CONFIG_FILE)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(contents))
	ConfigFile("test", string(contents)).Render(r.Context(), w)
}

func SetServerSettings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	s, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

	}
	Port, _ := strconv.ParseInt(r.Form.Get("vpn_port"), 10, 0)
	s.TCPPort = int(Port)
	WebServerPort, _ := strconv.ParseInt(r.Form.Get("web_port"), 10, 0)
	s.WebServerPort = int(WebServerPort)
	MaxConnections, _ := strconv.ParseInt(r.Form.Get("max_connections"), 10, 0)
	s.MaxConnections = int(MaxConnections)
	s.UseAsGateway = r.Form.Get("use_as_gateway") == "on"

	err = s.SaveSettings()
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
	s, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	s.DNSServer1 = r.Form.Get("dns1")
	s.DNSServer2 = r.Form.Get("dns2")
	s.PrivateAccess = r.Form.Get("private_access") == "on"

	err = s.SaveSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/settings/client", http.StatusSeeOther)
}

func SetAuthSettings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	settings, err := GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(r.Form.Get("strong_passwords"))
	max_auth_attempts, _ := strconv.ParseInt(r.Form.Get("max_auth_attempts"), 0, 0)
	settings.MaxAuthAttempts = int(max_auth_attempts)
	lockout_timeout, _ := time.ParseDuration(r.Form.Get("lockout_timeout"))
	settings.LockoutTimeout = lockout_timeout
	settings.EnforceStrongPW = r.Form.Get("strong_passwords") == "on"
	settings.AllowChangePW = r.Form.Get("change_password") == "on"
	err = settings.SaveSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/settings/auth", http.StatusSeeOther)
}

func SetConfigFileContent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	content := r.Form.Get("config")
	err = os.WriteFile(common.VPN_TCP_CONFIG_FILE, []byte(content), 755)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/settings/config", http.StatusSeeOther)
}
