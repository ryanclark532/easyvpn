package auth

import (
	"easyvpn/src/auth/auth_dtos"
	"easyvpn/src/database"
	"easyvpn/src/user"
	"easyvpn/src/utils"
	"fmt"
	"time"
)

func UserLogin(login *auth_dtos.LoginRequest) (*auth_dtos.LoginResponse, error) {
	var savedUser, err = user.GetUser(login.Username)
	if err != nil {
		return nil, err
	}

	if savedUser.Password != login.Password {
		return &auth_dtos.LoginResponse{
			Token:   "",
			IsAdmin: false,
			Error:   fmt.Sprintf("Password for %s is not correct", login.Username),
		}, nil
	}

	if !savedUser.Enabled {
		return &auth_dtos.LoginResponse{
			Token:   "",
			IsAdmin: false,
			Error:   fmt.Sprintf("User %s is not enabled", login.Username),
		}, nil
	}

	token, err := utils.CreateToken(savedUser)
	if err != nil {
		return nil, err
	}

	return &auth_dtos.LoginResponse{
		Token:           token,
		IsAdmin:         savedUser.IsAdmin,
		PasswordExpired: savedUser.PasswordExpiry.Before(time.Now()),
		ID:              savedUser.ID,
	}, nil
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

func ChangeUserPassword(changePassword *auth_dtos.ChangePasswordRequest) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE Users SET password='%s',password_expiry='%s'  WHERE ID=%s", changePassword.Password, time.Now().Add(30*24*time.Hour).Format(time.DateTime), changePassword.ID)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
