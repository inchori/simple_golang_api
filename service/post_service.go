package service

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/repository"
)

type IPostService interface {
	CreatePost(ctx context.Context, title, content string) (dto.PostResponse, error)
	GetPostByID(ctx context.Context, id int) (dto.PostResponse, error)
	DeleteByID(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, content, title string, id int) (dto.PostResponse, error)
}

type PostService struct {
	repo repository.IPostRepository
}

func NewPostService(postRepository repository.IPostRepository) IPostService {
	return &PostService{repo: postRepository}
}

func (p *PostService) CreatePost(ctx context.Context, title, content string) (dto.PostResponse, error) {
	post, err := p.repo.CreatePost(ctx, title, content)
	if err != nil {
		return dto.PostResponse{}, err
	}

	postResponse := dto.NewPostResponse(post)
	return postResponse, nil
}

func (p *PostService) GetPostByID(ctx context.Context, id int) (dto.PostResponse, error) {
	postByID, err := p.repo.GetPostByID(ctx, id)
	if err != nil {
		return dto.PostResponse{}, err
	}

	postResponse := dto.NewPostResponse(postByID)
	return postResponse, nil
}

func (p *PostService) DeleteByID(ctx context.Context, id int) error {
	return p.repo.DeleteByID(ctx, id)
}

func (p *PostService) UpdatePost(ctx context.Context, content, title string, id int) (dto.PostResponse, error) {
	updatePost, err := p.repo.UpdatePost(ctx, content, title, id)
	if err != nil {
		return dto.PostResponse{}, nil
	}

	postResponse := dto.NewPostResponse(updatePost)
	return postResponse, nil
}
