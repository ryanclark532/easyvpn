package user_dtos

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel  `bun:"table:users,alias:u"`
	ID             uint   `bun:",pk,autoincrement"`
	Name           string `bun:",notnull"`
	Username       string `bun:",notnull"`
	Password       string `bun:",notnull"`
	IsAdmin        bool
	Enabled        bool
	PasswordExpiry time.Time
}

type UserTest struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}
