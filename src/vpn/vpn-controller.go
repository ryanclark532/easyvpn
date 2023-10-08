package vpn

import (
	"easyvpn/src/utils"
	vpn_dtos "easyvpn/src/vpn/vpn-dtos"
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

func VpnOperationEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *vpn_dtos.VpnOperationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "VPNOperationEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = VpnOperation(req.Operation)
	if err != nil {
		utils.HandleError(err, "VPNOperationEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	GetServerStatusEndpoint(w, r)
}

func GetActiveConnectionsEndpoint(w http.ResponseWriter, r *http.Request) {

	response, err := GetActiveConnections()
	if err != nil {
		utils.HandleError(err, "VPNOperationEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["connections"] = response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "GetServerStatusEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
