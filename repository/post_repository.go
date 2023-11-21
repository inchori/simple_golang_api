package repository

import (
	"context"
	"grpc_identity/ent"
)

type IPostRepository interface {
	CreatePost(ctx context.Context, title, content string) (*ent.Post, error)
	GetPostByID(ctx context.Context, id int) (*ent.Post, error)
	DeleteByID(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, title, content string, id int) (*ent.Post, error)
}

type PostRepository struct {
	db *ent.PostClient
}

func (p *PostRepository) CreatePost(ctx context.Context, title, content string) (*ent.Post, error) {
	return p.db.
		Create().
		SetTitle(title).
		SetContent(content).
		Save(ctx)
}

func (p *PostRepository) GetPostByID(ctx context.Context, id int) (*ent.Post, error) {
	return p.db.Get(ctx, id)
}

func (p *PostRepository) DeleteByID(ctx context.Context, id int) error {
	return p.db.DeleteOneID(id).Exec(ctx)
}

func (p *PostRepository) UpdatePost(ctx context.Context, title, content string, id int) (*ent.Post, error) {
	return p.db.UpdateOneID(id).SetTitle(title).SetContent(content).Save(ctx)
}
