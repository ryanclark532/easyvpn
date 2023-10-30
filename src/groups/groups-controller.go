package groups

import (
	"easyvpn/src/groups/groups_dtos"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
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

func CreateGroupEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *groups_dtos.Group
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "CreateGroupEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = CreateGroup(req)
	if err != nil {
		utils.HandleError(err, "CreateGroupEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return

}

func CreateGroupMembershipEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *[]uint

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "CreateGroupMembersipEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = CreateGroupMembership(*req, chi.URLParam(r, "id"))
	if err != nil {
		utils.HandleError(err, "CreateGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
func DeleteGroupMembershipEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *[]uint

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "DeleteGroupMembershipEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DeleteGroupMembership(*req, chi.URLParam(r, "id"))
	if err != nil {
		utils.HandleError(err, "DeleteGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func DeleteGroupEndpoint(w http.ResponseWriter, r *http.Request) {
	err := DeleteGroup(chi.URLParam(r, "id"))
	if err != nil {
		utils.HandleError(err, "DeleteGroupEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func UpdateGroupEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *groups_dtos.Group
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(err, "UpdateGroupEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = UpdateGroup(req, chi.URLParam(r, "id"))
}
