package groups

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/groups/groups_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
	"fmt"
	"strconv"
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


func CreateGroup(group *groups_dtos.Group) (error) {
   _, err := database.DB.NewInsert().Model(group).Exec(context.Background())
   return err
}


func CreateGroupMembership(users []uint, groupid string) error{

    groupId, err := strconv.ParseUint(groupid, 10, 32)
    if err !=nil {
        return err
    }
    for _, e := range users {
        membership := groups_dtos.GroupMembership{
        UserID: e,
        GroupID: uint(groupId),
        }
     _, err =database.DB.NewInsert().Model(&membership).Exec(context.Background())
    }
return err
}


func DeleteGroupMembership(users []uint, groupid string) error{
    var err error
    for _, e:= range users {
        _, err= database.DB.NewDelete().Table("group_membership").Where("user_id = ? AND group_id = ?", e, groupid).Exec(context.Background())
    }
    return err
}
