package services

import (
	"easyvpn/src/database"
	"easyvpn/src/dtos"
)

func VerifyUser(username string, password string) (dtos.User, error) {
	loginRequest := dtos.LoginRequest{
		Username: username,
		Password: password,
	}

	var user, err = database.GetUser(loginRequest)
	if err != nil {
		return dtos.User{}, err
	}
	return user, nil
}
