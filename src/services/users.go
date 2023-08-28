package services

import (
	"easyvpn/src/dtos"
	"easyvpn/src/utils"
)

func VerifyUser(username string, password string) (*dtos.User, error) {
	db, err := utils.GetDB()
	if err != nil {
		return nil, utils.HandleError(err.Error(), "VerifyUser")
	}
	var user dtos.User
	result := db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
