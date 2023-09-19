package user_dtos

type FrontEndUser struct {
	ID       uint
	Username string
	Name     string
	IsAdmin  bool
	Enabled  bool
}

type FrontEndUsers struct {
	Users []FrontEndUser `json:"users"`
}
