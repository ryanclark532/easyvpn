package dtos

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
