package user

import (
	"context"
	"easyvpn/src/database"
	user_dtos "easyvpn/src/user/user-dtos"
	"easyvpn/src/utils"
	"fmt"
)

func GetUser(username string) (*user_dtos.User, error) {
	user := new(user_dtos.User)
	err := database.DB.NewSelect().Model(user).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers() (*[]user_dtos.User, error) {
	var users = new([]user_dtos.User)
	err := database.DB.NewSelect().Model(users).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *user_dtos.User) (*user_dtos.User, error) {
	_, err := database.DB.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return nil, err
	}

}

func DeleteUsers(IDs []int) error {

	tableName := "users"

	db, err := database.GetDB()
	if err != nil {
		return err
	}

	defer db.Close()
	query := fmt.Sprintf("DELETE FROM %s WHERE ID IN (%v)", tableName, utils.JoinInts(IDs, ", "))
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user user_dtos.FrontEndUser) (*user_dtos.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	query := fmt.Sprintf("UPDATE Users SET username='%s', name='%s', is_admin='%t', enabled='%t' WHERE id='%d'", user.Username, user.Name, user.IsAdmin, user.Enabled, user.ID)
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	u, err := GetUser(user.Username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func FormatUsers(users []user_dtos.User) []user_dtos.FrontEndUser {
	var formattedUsers []user_dtos.FrontEndUser
	for i := range users {
		u := users[i]
		formattedUsers = append(formattedUsers, FormatUser(&u))
	}
	return formattedUsers
}

func FormatUser(u *user_dtos.User) user_dtos.FrontEndUser {
	return user_dtos.FrontEndUser{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		IsAdmin:  u.IsAdmin,
		Enabled:  u.Enabled,
	}
}
