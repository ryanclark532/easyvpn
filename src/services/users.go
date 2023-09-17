package services

import (
	"database/sql"
	"easyvpn/src/database"
	"easyvpn/src/dtos"
	"easyvpn/src/utils"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type UsernameSearch struct {
	Username string
}

func VerifyUser(username string, password string) (*dtos.LoginResponse, error) {
	loginRequest := UsernameSearch{
		Username: username,
	}

	var user, err = GetUser(loginRequest)
	if err != nil {
		return nil, err
	}

	if utils.CheckEmpty(dtos.User{}, user) {
		return &dtos.LoginResponse{
			Token:   "",
			IsAdmin: true,
			Error:   fmt.Sprintf("User %s not found", username),
		}, nil
	}

	//TODO check password

	token, err := utils.CreateToken(user)
	if err != nil {

	}

	return &dtos.LoginResponse{
		Token:   token,
		IsAdmin: user.IsAdmin,
	}, nil
}

func GetUser(request interface{}) (dtos.User, error) {
	var user dtos.User
	db, err := database.GetDB()
	if err != nil {
		return dtos.User{}, err
	}
	requestType := reflect.TypeOf(request)
	var query string
	var values []interface{}

	query = "SELECT id, username, name, password, is_admin FROM Users WHERE "

	for i := 0; i < requestType.NumField(); i++ {
		field := requestType.Field(i)
		query += fmt.Sprintf("%s = ? AND ", field.Name)
		values = append(values, reflect.ValueOf(request).Field(i).Interface())
	}

	query = strings.TrimSuffix(query, "AND ")
	err = db.QueryRow(query, values...).Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin)
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
	defer rows.Close()

	var users []dtos.User

	for rows.Next() {
		var user dtos.User

		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin, &user.Enabled)
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
	query := "INSERT INTO Users (username, name, password, is_admin, enabled) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, Username, Name, Password, IsAdmin, Enabled)
	if err != nil {
		return dtos.User{}, err
	}
	return dtos.User{
		ID:       0,
		Username: Username,
		Name:     Name,
		Password: Password,
		IsAdmin:  IsAdmin,
		Enabled:  Enabled,
	}, nil
}

func DeleteUsers(IDs []int) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	tableName := "Users"

	query := fmt.Sprintf("DELETE FROM %s WHERE ID IN (%v)", tableName, utils.JoinInts(IDs, ", "))
	fmt.Println(query)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(request interface{}) (dtos.User, error) {
	db, err := database.GetDB()
	if err != nil {
		return dtos.User{}, err
	}

	requestType := reflect.TypeOf(request)
	tableName := "Users"
	primaryKeyField := "ID"

	query := fmt.Sprintf("UPDATE %s SET ", tableName)
	var setValues []interface{}

	for i := 0; i < requestType.Elem().NumField(); i++ {
		field := requestType.Elem().Field(i)
		fieldName := field.Name

		if fieldName == primaryKeyField {
			continue
		}

		query += fmt.Sprintf("%s = ?, ", fieldName)
		setValues = append(setValues, reflect.ValueOf(request).Elem().Field(i).Interface())
	}

	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf(" WHERE %s = ?", primaryKeyField)
	setValues = append(setValues, reflect.ValueOf(request).Elem().FieldByName(primaryKeyField).Interface())

	_, err = db.Exec(query, setValues...)
	if err != nil {
		return dtos.User{}, err
	}

	updatedUser := dtos.User{}
	selectQuery := fmt.Sprintf("SELECT id, username, name, password, is_admin FROM %s WHERE %s = ?", tableName, primaryKeyField)
	err = db.QueryRow(selectQuery, setValues[len(setValues)-1]).Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Name, &updatedUser.Password, &updatedUser.IsAdmin)
	if err != nil {
		return dtos.User{}, err
	}

	return updatedUser, nil
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
