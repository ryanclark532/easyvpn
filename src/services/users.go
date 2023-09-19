package services

import (
	"database/sql"
	"easyvpn/src/database"
	"easyvpn/src/dtos"
	"easyvpn/src/utils"
	"errors"
	"fmt"
	"time"
)

type UsernameSearch struct {
	Username string
}

type UpdateUserSearch struct {
	username string
	name     string
	id       uint
}

func VerifyUser(username string, password string) (*dtos.LoginResponse, error) {
	var user, err = GetUser(username)
	if err != nil {
		return nil, err
	}

	if utils.CheckEmpty(dtos.User{}, user) {
		return &dtos.LoginResponse{
			Token:   "",
			IsAdmin: false,
			Error:   fmt.Sprintf("User %s not found", username),
		}, nil
	}

	if user.Password != password {
		return &dtos.LoginResponse{
			Token:   "",
			IsAdmin: false,
			Error:   fmt.Sprintf("Password for %s is not correct", username),
		}, nil
	}

	if !user.Enabled {
		return &dtos.LoginResponse{
			Token:   "",
			IsAdmin: false,
			Error:   fmt.Sprintf("User %s is not enabled", username),
		}, nil
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.PasswordExpiry)

	return &dtos.LoginResponse{
		Token:           token,
		IsAdmin:         user.IsAdmin,
		PasswordExpired: user.PasswordExpiry.Before(time.Now()),
		ID:              user.ID,
	}, nil
}

func GetUser(username string) (dtos.User, error) {
	var user dtos.User
	db, err := database.GetDB()
	if err != nil {
		return dtos.User{}, err
	}
	query := fmt.Sprintf("SELECT id, username, name, password, is_admin, enabled, password_expiry FROM Users WHERE username ='%s'", username)

	err = db.QueryRow(query).Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin, &user.Enabled, &user.PasswordExpiry)
	if errors.Is(err, sql.ErrNoRows) {
		return dtos.User{}, nil
	}
	if err != nil {
		return dtos.User{}, err
	}

	return user, nil
}

func GetUsers() ([]dtos.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err // Return the error
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []dtos.User

	for rows.Next() {
		var user dtos.User

		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin, &user.Enabled, &user.PasswordExpiry)
		if err != nil {
			return nil, err // Return the error
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err // Return the error
	}

	return users, nil
}

func CreateUser(Username string, Name string, Password string, IsAdmin bool, Enabled bool) (dtos.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return dtos.User{}, err
	}
	query := "INSERT INTO Users (username, name, password, is_admin, enabled, password_expiry) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, Username, Name, Password, IsAdmin, Enabled, time.Now().Add(30*24*time.Hour))
	if err != nil {
		return dtos.User{}, err
	}

	user, err := GetUser(Username)
	if err != nil {
		return dtos.User{}, err
	}

	return user, nil
}

func DeleteUsers(IDs []int) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	tableName := "users"

	query := fmt.Sprintf("DELETE FROM %s WHERE ID IN (%v)", tableName, utils.JoinInts(IDs, ", "))
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user dtos.FrontEndUser) (*dtos.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("UPDATE Users SET username='%s', name='%s', is_admin='%t', enabled='%t' WHERE id='%d'", user.Username, user.Name, user.IsAdmin, user.Enabled, user.ID)
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	u, err := GetUser(user.Username)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func SetTempUserPassword(IDs []int) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	tempPassword := "changeme"

	query := fmt.Sprintf("UPDATE users SET password='%s',password_expiry='%s'  WHERE ID IN (%v)", tempPassword, time.Now().Format(time.DateTime), utils.JoinInts(IDs, ", "))
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func ChangeUserPassword(ID string, password string) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE Users SET password='%s',password_expiry='%s'  WHERE ID=%s", password, time.Now().Add(30*24*time.Hour).Format(time.DateTime), ID)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func FormatUsers(users []dtos.User) []dtos.FrontEndUser {
	var formattedUsers []dtos.FrontEndUser
	for i := range users {
		u := users[i]
		formattedUsers = append(formattedUsers, FormatUser(u))
	}
	return formattedUsers
}

func FormatUser(u dtos.User) dtos.FrontEndUser {
	return dtos.FrontEndUser{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		IsAdmin:  u.IsAdmin,
		Enabled:  u.Enabled,
	}
}
