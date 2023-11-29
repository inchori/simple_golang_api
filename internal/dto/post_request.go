package dto

type PostCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
