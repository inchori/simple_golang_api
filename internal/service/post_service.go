package service

import (
	"context"
	"github.com/inchori/grpc_identity/internal/ent"
	"github.com/inchori/grpc_identity/internal/repository"
)

type IPostService interface {
	CreatePost(ctx context.Context, title, content string, user *ent.User) (*ent.Post, error)
	GetPostByID(ctx context.Context, id int) (*ent.Post, error)
	GetPostByUserID(ctx context.Context, userID int) ([]*ent.Post, error)
	DeleteByID(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, content, title string, userID int) (*ent.Post, error)
}

type PostService struct {
	repo repository.IPostRepository
}

func NewPostService(postRepository repository.IPostRepository) IPostService {
	return &PostService{repo: postRepository}
}

func (p *PostService) CreatePost(ctx context.Context, title, content string, user *ent.User) (*ent.Post, error) {
	return p.repo.CreatePost(ctx, title, content, user)
}

func (p *PostService) GetPostByID(ctx context.Context, id int) (*ent.Post, error) {
	return p.repo.GetPostByID(ctx, id)
}

func (p *PostService) GetPostByUserID(ctx context.Context, userID int) ([]*ent.Post, error) {
	return p.repo.GetPostByUserID(ctx, userID)
}

func (p *PostService) DeleteByID(ctx context.Context, id int) error {
	return p.repo.DeleteByID(ctx, id)
}

func (p *PostService) UpdatePost(ctx context.Context, content, title string, id int) (*ent.Post, error) {
	return p.repo.UpdatePost(ctx, content, title, id)
}
