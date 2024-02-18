package groups

import (
	"context"
	"easyvpn/internal/database"
	"easyvpn/internal/user"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

type GroupWithMembership struct {
	Group
	members []user.User
}

type GroupMembership struct {
	bun.BaseModel `bun:"table:group_membership,alias:gm"`
	ID            int `bun:",pk,autoincrement" json:"id"`
	UserID        int `bun:",notnull" json:"user_id"`
	GroupID       int `bun:",notnull" json:"group"`
}

type Group struct {
	bun.BaseModel `bun:"table:groups,alias:g"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name"`
	MemberCount   int    `json:"member_count" bun:"-"`
	IsAdmin       bool   `bun:",notnull" json:"is_admin"`
	Roles         string `bun:",notnull" json:"roles"`
	Enabled       bool   `bun:",notnull" json:"enabled"`
}

func (g Group) GetMemberships() ([]user.User, error) {
	var users []user.User
	err := database.DB.NewSelect().
		Model(&users).
		Join("INNER JOIN group_membership gm ON u.id = gm.user_id").
		Where("gm.group_id = ?", g.ID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return users, nil
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
		members, err := group.GetMemberships()
		if err != nil {
			continue
		}
		x := GroupWithMembership{
			Group:   group,
			members: members,
		}
		groupsWithMembership = append(groupsWithMembership, x)
	}
	for _, x := range groupsWithMembership {
		fmt.Println(x.members)
	}

	Groups("hello", groupsWithMembership, *users, chi.URLParam(r, "username"), user.CompleteRoles).Render(r.Context(), w)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	group := Group{
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
			groupMembership := GroupMembership{
				UserID:  int(userId),
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
		members, err := group.GetMemberships()
		if err != nil {
			continue
		}
		x := GroupWithMembership{
			Group:   group,
			members: members,
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

	_, err = database.DB.NewDelete().Model((*GroupMembership)(nil)).Where("group_id = ?", groupId).Exec(context.Background())
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

		groupMembership := GroupMembership{
			UserID:  int(userId),
			GroupID: int(groupId),
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

	group := Group{
		Name:        r.Form.Get("name"),
		MemberCount: 0,
		IsAdmin:     r.Form.Get("admin") == "on",
		Enabled:     r.Form.Get("enabled") == "on",
		Roles:       strings.Join(r.Form["roles"], ","),
	}
	_, err = database.DB.NewUpdate().Model(group).Where("id = ?", groupId).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/groups", http.StatusSeeOther)
}

func GroupContainsMember(s []user.User, e user.User) bool {
	for _, value := range s {
		if value.ID == e.ID {
			return true
		}
	}
	return false
}

func GetGroups() ([]Group, error) {
	groups := new([]Group)
	err := database.DB.NewSelect().Model(groups).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	var g []Group
	for _, group := range *groups {
		m := new([]GroupMembership)
		err = database.DB.NewSelect().Model(m).Where("group_id = ?", group.ID).Scan(context.Background())
		group.MemberCount = len(*m)
		g = append(g, group)
	}
	return g, err
}
