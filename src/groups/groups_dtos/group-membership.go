package groups_dtos

import (

	"github.com/uptrace/bun"
)

type GroupMembership struct {
    bun.BaseModel `bun:"table:group_membership,alias:gm"`
    ID            uint            `bun:",pk,autoincrement" json:"id"`
    UserID          uint `bun:",notnull" json:"user_id"`
    GroupID          uint `bun:",notnull" json:"group"`
}
