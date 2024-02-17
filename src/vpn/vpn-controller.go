package vpn

import (
	"bufio"
	"easyvpn/src/common"
	"easyvpn/src/settings"
	"easyvpn/src/utils"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Log struct {
	LogTime time.Time
	LogText string
}

type ServerOverview struct {
	Name           string
	IP             string
	MaxConnections int
	ActiveUsers    int
	VpnPort        int
	WebPort        int
}
type ServerConnection struct {
	CommonName     string
	Address        string
	BytesRec       string
	BytesSent      string
	ConnectedSince time.Time
}

func GetActiveUsersPage(w http.ResponseWriter, r *http.Request) {
	activeUsers, err := GetActiveConnections(r.URL.Query().Get("search"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if activeUsers == nil {
		var x []ServerConnection
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

func StatusOverviewPage(w http.ResponseWriter, r *http.Request) {
	status, err := utils.GetVpnServerStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var color string
	if status == "running" {
		color = "green"
	} else if status == "notRunning" {
		color = "red"
	} else {
		color = "yellow"
	}
	s, err := settings.GetSettings()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	currentConnections, err := GetActiveConnections("")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name, _ := os.Hostname()
	current := ServerOverview{
		Name:           name,
		IP:             s.IPAddress,
		MaxConnections: s.MaxConnections,
		ActiveUsers:    len(*currentConnections),
		VpnPort:        s.TCPPort,
		WebPort:        s.WebServerPort,
	}

	StatusOverview("test", current, status, color).Render(r.Context(), w)
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

func VpnOperation(w http.ResponseWriter, r *http.Request) {
	operation := r.URL.Query().Get("operation")
	switch operation {
	case "start":
		fmt.Println("sttart")
		utils.StartVPNServer()
	case "stop":
		fmt.Println("stop")
		utils.StopVPNServer()
	case "restart":
		fmt.Println("restart")
		utils.StopVPNServer()
		utils.StartVPNServer()
	}
}

func GetActiveConnections(searchterm string) (*[]ServerConnection, error) {
	conn, err := utils.ConnectTelnet("localhost:7505")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = utils.CommandTelnet("status", conn)
	if err != nil {
		return nil, err
	}
	out, err := utils.ReadTelnet(conn)
	if err != nil {
		return nil, err
	}
	outString := strings.Join(out, "")
	return formatServerConnection(outString, searchterm)
}

func formatServerConnection(output string, searchterm string) (*[]ServerConnection, error) {
	lines := strings.Split(output, "\r\n")

	headerLine := ""
	for i, line := range lines {
		if line == "Common Name,Real Address,Bytes Received,Bytes Sent,Connected Since" {
			headerLine = lines[i]
			break
		}
	}

	headerIndex := strings.Index(output, headerLine)

	dataLines := strings.Split(output[headerIndex:], "\r\n")[1:]

	var connections []ServerConnection

	for _, dataLine := range dataLines {
		fields := strings.Split(dataLine, ",")
		if len(fields) == 5 {
			commonName := fields[0]
			realAddress := fields[1]
			bytesReceived := fields[2]
			bytesSent := fields[3]
			connectedSince := fields[4]
			parsedTime, err := time.Parse(time.DateTime, connectedSince)
			if err != nil {
				return nil, err
			}

			y := ServerConnection{
				CommonName:     commonName,
				Address:        realAddress,
				BytesRec:       bytesReceived,
				BytesSent:      bytesSent,
				ConnectedSince: parsedTime,
			}
			if fuzzy.Match(searchterm, commonName) {
				connections = append(connections, y)
			}
		}
	}

	return &connections, nil

}
