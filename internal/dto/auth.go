package dto

type CreateUserTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserTokenResponse struct {
	Token string `json:"token"`
}
