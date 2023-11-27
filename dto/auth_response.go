package dto

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}
