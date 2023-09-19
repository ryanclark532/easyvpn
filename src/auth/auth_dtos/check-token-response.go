package auth_dtos

type CheckTokenResponse struct {
	IsAdmin         bool `json:"is_admin"`
	TokenValid      bool `json:"token_valid"`
	PasswordExpired bool `json:"password_expired"`
}
