package user

type RegisterRequest struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type LoginRequest struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
}
