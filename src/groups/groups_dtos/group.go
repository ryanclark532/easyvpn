package groups_dtos

import "github.com/uptrace/bun"

type Group struct {
	bun.BaseModel `bun:"table:groups,alias:g"`
	ID            uint   `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name"`
	MemberCount   int    `json:"member_count" bun:"-"`
	IsAdmin       bool   `bun:",notnull" json:"is_admin"`
	Roles         string `bun:",notnull" json:"roles"`
	Enabled       bool   `bun:",notnull" json:"enabled"`
}
