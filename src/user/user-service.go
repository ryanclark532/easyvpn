package user

import (
	"context"
	"easyvpn/src/database"
	user_dtos "easyvpn/src/user/user-dtos"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(username string) (*user_dtos.User, error) {
	user := new(user_dtos.User)
	err := database.DB.NewSelect().Model(user).Where("username = ?", username).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers() (*[]user_dtos.User, error) {
	var users = new([]user_dtos.User)
	err := database.DB.NewSelect().Model(users).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *user_dtos.User) (*user_dtos.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.PasswordExpiry = time.Now().Add(time.Hour * 1000)
	_, err = database.DB.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id string) error {
	_, err := database.DB.NewDelete().Model((*user_dtos.User)(nil)).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *user_dtos.User, id string) error {
	_, err := database.DB.NewUpdate().Model((*user_dtos.User)(nil)).Set("username = ?, name = ?, is_admin = ?, enabled = ?", user.Username, user.Name, user.IsAdmin, user.Enabled).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func SetPassword(userId string, ps *user_dtos.PasswordResetRequest) error{
	if ps.Password != ps.Confirm {
		return fmt.Errorf("Password and Confirmation do not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(ps.Password), 10)
	if err != nil {
		return err
	}

	_, err =database.DB.NewUpdate().Model((*user_dtos.User)(nil)).Set("password = ?, password_expiry = ?", hash, time.Now().Format(time.DateTime)).Where("id = ?", userId).Exec(context.Background())
	return err	
}

func FormatUsers(users []user_dtos.User) []user_dtos.FrontEndUser {
	var formattedUsers []user_dtos.FrontEndUser
	for i := range users {
		u := users[i]
		formattedUsers = append(formattedUsers, FormatUser(u))
	}
	return formattedUsers
}

func FormatUser(u user_dtos.User) user_dtos.FrontEndUser {
	return user_dtos.FrontEndUser{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		IsAdmin:  u.IsAdmin,
		Enabled:  u.Enabled,
	}
}

func AuthUser(username string, password string) (bool, error) {
	user, err := GetUser(username)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
