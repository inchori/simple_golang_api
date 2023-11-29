package repository

import (
	"context"
	"github.com/inchori/grpc_identity/internal/ent"
	"github.com/inchori/grpc_identity/internal/ent/post"
	"github.com/inchori/grpc_identity/internal/ent/user"
)

type IPostRepository interface {
	CreatePost(ctx context.Context, title, content string, user *ent.User) (*ent.Post, error)
	GetPostByID(ctx context.Context, id int) (*ent.Post, error)
	GetPostByUserID(ctx context.Context, userID int) ([]*ent.Post, error)
	DeleteByID(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, title, content string, id int) (*ent.Post, error)
}

type PostRepository struct {
	db *ent.PostClient
}

func NewPostRepository(db *ent.PostClient) IPostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) CreatePost(ctx context.Context, title, content string, user *ent.User) (*ent.Post, error) {
	return p.db.
		Create().
		SetTitle(title).
		SetContent(content).
		SetUser(user).
		Save(ctx)
}

func (p *PostRepository) GetPostByID(ctx context.Context, id int) (*ent.Post, error) {
	return p.db.Get(ctx, id)
}

func (p *PostRepository) GetPostByUserID(ctx context.Context, userID int) ([]*ent.Post, error) {
	return p.db.Query().Where(post.HasUserWith(user.IDEQ(userID))).WithUser().All(ctx)
}

func (p *PostRepository) DeleteByID(ctx context.Context, id int) error {
	return p.db.DeleteOneID(id).Exec(ctx)
}

func (p *PostRepository) UpdatePost(ctx context.Context, title, content string, id int) (*ent.Post, error) {
	return p.db.UpdateOneID(id).SetTitle(title).SetContent(content).Save(ctx)
}
