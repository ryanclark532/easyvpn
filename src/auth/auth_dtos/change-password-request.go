package auth_dtos

type ChangePasswordRequest struct {
	Password string `json:"password"`
	ID       string `json:"id"`
}
