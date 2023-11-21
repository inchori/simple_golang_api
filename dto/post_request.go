package dto

type PostCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostUpdateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
