package user_dtos

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel  `bun:"table:users,alias:u"`
	ID             uint      `bun:",pk,autoincrement"`
	Name           string    `bun:",notnull" json:"name"`
	Username       string    `bun:",notnull" json:"username"`
	Password       string    `bun:",notnull" json:"password"`
	IsAdmin        bool      `json:"is_admin"`
	Enabled        bool      `json:"enabled"`
	PasswordExpiry time.Time `json:"password_expiry"`
}

type UserTest struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}
