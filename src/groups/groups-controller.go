package groups

import (
	"easyvpn/src/groups/groups_dtos"
	"easyvpn/src/utils"
	"encoding/json"
	"net/http"
)

func GetGroupsEndpoint(w http.ResponseWriter, r *http.Request) {
	response, err := GetGroups()
	if err != nil {
		utils.HandleError(err, "GetGroupsEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["groups"] = response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "GetServerStatusEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetGroupMembershipEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *groups_dtos.Group
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "CreateUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := GetMemberships(req.ID)
	if err != nil {
		utils.HandleError(err, "GetGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{}
	responseData["members"] = response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		utils.HandleError(err, "GetGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
