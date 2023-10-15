package groups

import (
	"context"
	"easyvpn/src/database"
	"easyvpn/src/groups/groups_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
)

func GetGroups() (*[]groups_dtos.Group, error) {
	var groups = new([]groups_dtos.Group)
	err := database.DB.NewSelect().Model(groups).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func GetMemberships(groupId string) (*[]user_dtos.User, error) {

	var memberships = new([]groups_dtos.GroupMembership)
	err := database.DB.NewSelect().Model(memberships).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	var users []user_dtos.User

	for _, membership := range *memberships {
		if membership.User != nil {
			users = append(users, *membership.User)
		}
	}

	return &users, nil
}
