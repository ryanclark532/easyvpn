package services

import (
	"easyvpn/src/database"
	"easyvpn/src/dtos"
	"fmt"
	"reflect"
	"strings"
)

func VerifyUser(username string, password string) (dtos.User, error) {
	loginRequest := dtos.LoginRequest{
		Username: username,
		Password: password,
	}

	var user, err = GetUser(loginRequest)
	if err != nil {
		return dtos.User{}, err
	}
	return user, nil
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

	query = "SELECT ID, username, name, password, IsAdmin FROM users WHERE "

	for i := 0; i < requestType.NumField(); i++ {
		field := requestType.Field(i)
		query += fmt.Sprintf("%s = ? AND ", field.Name)
		values = append(values, reflect.ValueOf(request).Field(i).Interface())
	}

	query = strings.TrimSuffix(query, "AND ")

	// Execute the query
	err = db.QueryRow(query, values...).Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin)
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
	fmt.Println(db)

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err // Return the error
	}
	defer rows.Close()

	var users []dtos.User

	for rows.Next() {
		var user dtos.User

		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.IsAdmin)
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
	query := "INSERT INTO Users (username, name, password, IsAdmin, enabled) VALUES (?, ?, ?, ?, ?)"
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
	}, nil
}
