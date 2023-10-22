package groups

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/groups/groups_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
	"fmt"
)

func GetGroups() (*[]groups_dtos.Group, error) {
	var groups = new([]groups_dtos.Group)
	err := database.DB.NewSelect().Model(groups).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func GetMembershipsForGroup(groupId string) (*[]user_dtos.User, error) {
    fmt.Println(groupId)
    users:= new([]user_dtos.User)
    err := database.DB.NewSelect().
    Model(users).
    Join("INNER JOIN group_membership gm ON u.id = gm.user_id").
    Where("gm.group_id = ?", groupId).
    Scan(context.Background())
    if err != nil {
        return nil, err
    }

    fmt.Println(users)

    return users, nil
}
func GetMembershipsforUser(userId uint) (*[]groups_dtos.Group, error) {
	var memberships = new([]groups_dtos.GroupMembership)
	err := database.DB.NewSelect().Model(memberships).Relation("Group").Where("user.id = ?", userId).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	var groups []groups_dtos.Group

    return &groups, nil
}
