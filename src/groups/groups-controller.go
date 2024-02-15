package groups

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/groups/groups_dtos"
	"easyvpn/src/logging"
	"easyvpn/src/user"
	user_dtos "easyvpn/src/user/user-dtos"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type GroupWithMembership struct {
	groups_dtos.Group
	members []user_dtos.User
}

func GetGroupsEndpoint(w http.ResponseWriter, r *http.Request) {
	response, err := GetGroups()
	if err != nil {
		logging.HandleError(err, "GetGroupsEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logging.HandleError(err, "GetGroupsEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetGroupMembershipEndpoint(w http.ResponseWriter, r *http.Request) {
	response, err := GetMembershipsForGroup(chi.URLParam(r, "id"))
	if err != nil {
		logging.HandleError(err, "GetGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logging.HandleError(err, "GetGroupMembershipEndpoint")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateGroupMembershipEndpoint(w http.ResponseWriter, r *http.Request) {
	var req *[]uint

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.HandleError(err, "CreateGroupMembersipEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = CreateGroupMembership(*req, chi.URLParam(r, "id"))
	if err != nil {
		logging.HandleError(err, "CreateGroupMembershipEndpoint")
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
		logging.HandleError(err, "DeleteGroupMembershipEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DeleteGroupMembership(*req, chi.URLParam(r, "id"))
	if err != nil {
		logging.HandleError(err, "DeleteGroupMembershipEndpoint")
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
		logging.HandleError(err, "UpdateGroupEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = UpdateGroup(req, chi.URLParam(r, "id"))
	if err != nil {
		logging.HandleError(err, "UpdateGroupEndpoint")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GroupsPage(w http.ResponseWriter, r *http.Request) {

	groups, err := GetGroups()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	users, err := user.GetUsers("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var groupsWithMembership []GroupWithMembership

	for _, group := range groups {
		members, err := GetMembershipsForGroup(strconv.Itoa(int(group.ID)))
		if err != nil {
			continue
		}
		x := GroupWithMembership{
			Group:   group,
			members: *members,
		}
		groupsWithMembership = append(groupsWithMembership, x)
	}
	for _, x := range groupsWithMembership {
		fmt.Println(x.members)
	}

	Groups("hello", groupsWithMembership, *users, chi.URLParam(r, "username"), user_dtos.CompleteRoles).Render(r.Context(), w)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	group := groups_dtos.Group{
		Name:        r.Form.Get("name"),
		MemberCount: 0,
		IsAdmin:     r.Form.Get("admin") == "on",
		Enabled:     r.Form.Get("enabled") == "on",
		Roles:       strings.Join(r.Form["roles"], ","),
	}

	err = database.DB.NewInsert().Model(&group).Scan(context.Background(), &group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	formObjs := strings.Split(r.Form.Encode(), "&")
	var addedUsers int
	for _, v := range formObjs {
		if strings.HasSuffix(v, "-group=on") {
			userId, err := strconv.ParseInt(strings.TrimSuffix(v, "-group=on"), 36, 64)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			groupMembership := groups_dtos.GroupMembership{
				UserID:  uint(userId),
				GroupID: group.ID,
			}
			_, err = database.DB.NewInsert().Model(&groupMembership).Exec(context.Background())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			addedUsers++

		}
	}
	group.MemberCount = addedUsers
	_, err = database.DB.NewUpdate().Model(&group).Where("ID = ?", group.ID).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/groups", http.StatusSeeOther)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	_, err := database.DB.NewDelete().Table("groups").Where("ID = ?", chi.URLParam(r, "id")).Exec(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	groups, err := GetGroups()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	users, err := user.GetUsers("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var groupsWithMembership []GroupWithMembership

	for _, group := range groups {
		members, err := GetMembershipsForGroup(strconv.Itoa(int(group.ID)))
		if err != nil {
			continue
		}
		x := GroupWithMembership{
			Group:   group,
			members: *members,
		}
		groupsWithMembership = append(groupsWithMembership, x)
	}

	GroupsTable(groupsWithMembership, *users, "").Render(r.Context(), w)
}

func UpdateGroupPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(r.Form.Encode())

	groupId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = database.DB.NewDelete().Model((*groups_dtos.GroupMembership)(nil)).Where("group_id = ?", groupId).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	formObjs := strings.Split(r.Form.Encode(), "&")
	for _, v := range formObjs {
		if !strings.HasSuffix(v, "-group=on") {
			continue
		}
		userId, err := strconv.ParseInt(strings.TrimSuffix(v, "-group=on"), 36, 64)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		groupMembership := groups_dtos.GroupMembership{
			UserID:  uint(userId),
			GroupID: uint(groupId),
		}
		_, err = database.DB.NewInsert().Model(&groupMembership).Ignore().Exec(context.Background())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}

	group := groups_dtos.Group{
		Name:        r.Form.Get("name"),
		MemberCount: 0,
		IsAdmin:     r.Form.Get("admin") == "on",
		Enabled:     r.Form.Get("enabled") == "on",
		Roles:       strings.Join(r.Form["roles"], ","),
	}
	err = UpdateGroup(&group, strconv.Itoa(int(groupId)))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/groups", http.StatusSeeOther)
}

func GroupContainsMember(s []user_dtos.User, e user_dtos.User) bool {
	for _, value := range s {
		if value.ID == e.ID {
			return true
		}
	}
	return false
}
