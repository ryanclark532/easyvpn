package user_dtos

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel  `bun:"table:users,alias:u"`
	ID             uint      `bun:",pk,autoincrement" json:"id"`
	Name           string    `bun:",notnull" json:"name"`
	Username       string    `bun:",notnull" json:"username"`
	Password       string    `bun:",notnull" json:"password"`
	Roles          string    `bun:",notnull" json:"roles"`
	IsAdmin        bool      `json:"is_admin"`
	Enabled        bool      `json:"enabled"`
	PasswordExpiry time.Time `json:"password_expiry"`
}
