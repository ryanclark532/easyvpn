package groups

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/groups/groups_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
	"strconv"
)

func GetGroups() (*[]groups_dtos.Group, error) {
	groups := new([]groups_dtos.Group)
	err := database.DB.NewSelect().Model(groups).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	var g []groups_dtos.Group
	for _, group := range *groups{
		m := new([]groups_dtos.GroupMembership)
		err = database.DB.NewSelect().Model(m).Where("group_id = ?", group.ID).Scan(context.Background())
		group.MemberCount = len(*m)
		g = append(g, group)
	}	
	return &g, err
}

func GetMembershipsForGroup(groupId string) (*[]user_dtos.User, error) {
	users := new([]user_dtos.User)
	err := database.DB.NewSelect().
		Model(users).
		Join("INNER JOIN group_membership gm ON u.id = gm.user_id").
		Where("gm.group_id = ?", groupId).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return users, nil
}
func GetMembershipsforUser(userId uint) ([]groups_dtos.Group, error) {
	memberships := new([]groups_dtos.GroupMembership)
	err := database.DB.NewSelect().Model(memberships).Relation("Group").Where("user.id = ?", userId).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	var groups []groups_dtos.Group
	return groups, nil
}


func CreateGroupMembership(users []uint, groupid string) error {
	groupId, err := strconv.ParseUint(groupid, 10, 32)
	if err != nil {
		return err
	}
	for _, e := range users {
		membership := groups_dtos.GroupMembership{
			UserID:  e,
			GroupID: uint(groupId),
		}
		_, err = database.DB.NewInsert().Model(&membership).Exec(context.Background())
	}

	_, err = database.DB.NewUpdate().Table("groups").Set("member_count = member_count + ?", len(users)).Where("id = ?", groupid).Exec(context.Background())
	return err
}

func DeleteGroupMembership(users []uint, groupid string) error {
	var err error
	for _, e := range users {
		_, err = database.DB.NewDelete().Table("group_membership").Where("user_id = ? AND group_id = ?", e, groupid).Exec(context.Background())
	}

	_, err = database.DB.NewUpdate().Table("groups").Set("member_count = member_count - ?", len(users)).Where("id = ?", groupid).Exec(context.Background())
	return err
}


func UpdateGroup(group *groups_dtos.Group, groupid string) error {
	_, err := database.DB.NewUpdate().Model(group).
		Set("name = ?, is_admin = ?, enabled = ?", group.Name, group.IsAdmin, group.Enabled).
		Where("id = ?", groupid).Exec(context.Background())
	return err
}