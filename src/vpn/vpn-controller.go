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
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lithammer/fuzzysearch/fuzzy"
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
	response, err := GetActiveConnections("blah")
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
	activeUsers, err := GetActiveConnections(r.URL.Query().Get("search"))
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
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 10)
	if err != nil {
		page = 1
	}
	logs, err := GetVPNLogs(int(page), r.URL.Query().Get("search"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(r.URL.Query().Get("search"))
	Logs("test", logs, int(page), r.URL.Query().Get("search"), len(logs) == 100, int(page) != 1).Render(r.Context(), w)
}

func DisconnectClient(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	conn, err := utils.ConnectTelnet("localhost:7505")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	err = utils.CommandTelnet(fmt.Sprintf("kill %s", username), conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = utils.ReadTelnet(conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/vpn/active-connections", http.StatusSeeOther)
}

func GetVPNLogs(page int, searchterm string) ([]Log, error) {
	file, err := os.Open(common.VPNLOG_FILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var logs []Log
	scanner := bufio.NewScanner(file)

	if page > 1 {
		for i := 0; i < page*100 && scanner.Scan(); i++ {
			continue
		}
	}

	for i := 0; i < 100 && scanner.Scan(); {
		line := scanner.Text()
		log, _ := parseLogLine(line)
		if (searchterm == "" || fuzzy.Match(searchterm, log.LogText)) && log.LogText != "" {
			logs = append(logs, log)
			i++
		}
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].LogTime.After(logs[j].LogTime)
	})
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
