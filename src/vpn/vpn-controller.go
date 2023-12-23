package vpn

import (
	"easyvpn/src/logging"
	"easyvpn/src/utils"
	vpn_dtos "easyvpn/src/vpn/vpn-dtos"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

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

func GetVpnLogs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	err = watcher.Add(`C:\Program Files\OpenVPN\log\server-dev.log`)
	if err != nil {
		fmt.Println(err)
	}

	content, _ := os.ReadFile(`C:\Program Files\OpenVPN\log\server-dev.log`)
	conn.WriteMessage(1, content)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			content, _ := os.ReadFile(`C:\Program Files\OpenVPN\log\server-dev.log`)
			conn.WriteMessage(1, content)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
