package repository

import (
	"context"
	"grpc_identity/ent"
	"grpc_identity/ent/user"
	"time"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, name, email, password string) (*ent.User, error)
	GetUserByID(ctx context.Context, id int) (*ent.User, error)
	GetUserByName(ctx context.Context, name string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, name, password string, id int) (*ent.User, error)
}

type UserRepository struct {
	db *ent.UserClient
}

func NewUserRepository(db *ent.UserClient) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, name, email, password string) (*ent.User, error) {
	return u.db.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
}

func (u *UserRepository) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	return u.db.Query().Where(user.IDEQ(id)).WithPosts().Only(ctx)
}

func (u *UserRepository) GetUserByName(ctx context.Context, name string) (*ent.User, error) {
	return u.db.Query().Where(user.Name(name)).Only(ctx)
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.db.Query().Where(user.Email(email)).Only(ctx)
}

func (u *UserRepository) DeleteByID(ctx context.Context, id int) error {
	return u.db.DeleteOneID(id).Exec(ctx)
}

func (u *UserRepository) UpdateUser(ctx context.Context, name, password string, id int) (*ent.User, error) {
	return u.db.
		UpdateOneID(id).
		SetName(name).
		SetPassword(password).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}
