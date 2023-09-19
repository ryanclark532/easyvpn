package auth_dtos

type LoginResponse struct {
	Token           string `json:"token"`
	IsAdmin         bool   `json:"is_admin"`
	Error           string `json:"error"`
	PasswordExpired bool   `json:"password_expired"`
	ID              uint   `json:"id"`
}
