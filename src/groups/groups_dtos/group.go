package groups_dtos

type Group struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MemberCount int    `json:"member_count"`
	IsAdmin     bool   `json:"is_admin"`
	Enabled     bool   `json:"enabled"`
}
