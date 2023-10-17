package groups_dtos

import (
	user_dtos "easyvpn/src/user/user-dtos"

	"github.com/uptrace/bun"
)

type GroupMembership struct {
	bun.BaseModel `bun:"table:group_membership,alias:gm"`
	ID            uint            `bun:",pk,autoincrement" json:"id"`
	User          *user_dtos.User `bun:"rel:has-one,join:id=id"`
	Group         *Group          `bun:"rel:has-one,join:id=id"`
}
