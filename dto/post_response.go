package dto

import "grpc_identity/ent"

type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPostResponse(post *ent.Post) PostResponse {
	return PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}

type PostsResponse struct {
	Count int            `json:"count"`
	Posts []PostResponse `json:"posts"`
}

func NewPostsResponse(postsResponse []PostResponse) PostsResponse {
	return PostsResponse{
		Count: len(postsResponse),
		Posts: postsResponse,
	}
}
