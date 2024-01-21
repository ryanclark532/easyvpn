package vpn

import (
	"bufio"
	"easyvpn/src/common"
	"easyvpn/src/logging"
	"easyvpn/src/utils"
	vpn_dtos "easyvpn/src/vpn/vpn-dtos"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Log struct {
	LogTime time.Time
	LogText string
}

func GetServerStatusEndpoint(w http.ResponseWriter, r *http.Request) {
	status, err := utils.GetVpnServerStatus()
	if err != nil {
		logging.HandleError(err, "GetServerStatusEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["status"] = status

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		logging.HandleError(err, "GetServerStatusEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func VpnOperationEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *vpn_dtos.VpnOperationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.HandleError(err, "VPNOperationEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = VpnOperation(req.Operation)
	if err != nil {
		logging.HandleError(err, "VPNOperationEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	GetServerStatusEndpoint(w, r)
}

func GetActiveConnectionsEndpoint(w http.ResponseWriter, r *http.Request) {
	response, err := GetActiveConnections()
	if err != nil {
		logging.HandleError(err, "GetActiveConnectionsEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["connections"] = response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		logging.HandleError(err, "GetActiveConnectionsEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetActiveUsersPage(w http.ResponseWriter, r *http.Request) {
	activeUsers, err := GetActiveConnections()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if activeUsers == nil {
		var x []vpn_dtos.ServerConnection
		p := &x
		ActiveUsers("test", p, "").Render(r.Context(), w)
		return
	}
	ActiveUsers("test", activeUsers, "").Render(r.Context(), w)
}

func GetVpnLogsPage(w http.ResponseWriter, r *http.Request) {

	logs, err := GetVPNLogs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Logs("test", logs, "").Render(r.Context(), w)
}

func GetVPNLogs() ([]Log, error) {
	file, err := os.Open(common.VPNLOG_FILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var logs []Log
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log, _ := parseLogLine(line)
		if log.LogText != "" {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

func parseLogLine(line string) (Log, error) {
	parts := strings.SplitN(line, " ", 4)

	if len(parts) < 4 {
		return Log{}, fmt.Errorf("Invalid log format: %s", line)
	}

	logTime, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
	if err != nil {
		return Log{}, fmt.Errorf("Error parsing timestamp: %s", err)
	}

	log := Log{
		LogTime: logTime,
		LogText: parts[3],
	}

	return log, nil
}
