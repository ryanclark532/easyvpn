package user

import (
	"database/sql"
	"easyvpn/src/database"
	user_dtos "easyvpn/src/user/user-dtos"
	"easyvpn/src/utils"
	"errors"
	"fmt"
	"time"
)

func GetUser(username string) (*user_dtos.User, error) {
	var user user_dtos.User

	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT id, username, name, password, is_admin, enabled, password_expiry FROM Users WHERE username ='%s'", username)
	err = db.QueryRow(query).Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin, &user.Enabled, &user.PasswordExpiry)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers() (*[]user_dtos.User, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []user_dtos.User

	for rows.Next() {
		var user user_dtos.User

		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin, &user.Enabled, &user.PasswordExpiry)
		if err != nil {
			return nil, err // Return the error
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err // Return the error
	}

	return &users, nil
}

func CreateUser(user *user_dtos.CreateUserRequest) (*user_dtos.User, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	query := "INSERT INTO Users (username, name, password, is_admin, enabled, password_expiry) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, user.Username, user.Name, user.Password, user.IsAdmin, user.Enabled, time.Now().Add(30*24*time.Hour))
	if err != nil {
		return nil, err
	}

	modifiedUser, err := GetUser(user.Username)
	if err != nil {
		return nil, err
	}

	return modifiedUser, nil
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
