package user_dtos

import "time"

type User struct {
	ID             uint
	Username       string
	Name           string
	Password       string
	IsAdmin        bool
	Enabled        bool
	PasswordExpiry time.Time
}
