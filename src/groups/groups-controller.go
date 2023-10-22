package groups

import (
	"easyvpn/src/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	response, err := GetMembershipsForGroup(chi.URLParam(r, "id"))
	if err != nil {
		utils.HandleError(err, "GetGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.HandleError(err, "GetGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
