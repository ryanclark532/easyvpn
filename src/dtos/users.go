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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CheckTokenRequest struct {
	Token string `json:"token"`
}

type CheckTokenResponse struct {
	IsAdmin    bool `json:"is_admin"`
	TokenValid bool `json:"token_valid"`
}
