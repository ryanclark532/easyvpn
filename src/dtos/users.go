package dtos

type User struct {
	ID       uint
	Username string
	Name     string
	Password string
	IsAdmin  bool
	Enabled  bool
}

type CreateUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	Enabled  bool   `json:"enabled"`
}

type UserID struct {
	ID []int `json:"id"`
}

type FrontEndUser struct {
	ID       uint
	Username string
	Name     string
	IsAdmin  bool
	Enabled  bool
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
	Error   string `json:"error"`
}

type CheckTokenRequest struct {
	Token string `json:"token"`
}

type CheckTokenResponse struct {
	IsAdmin    bool `json:"is_admin"`
	TokenValid bool `json:"token_valid"`
}
