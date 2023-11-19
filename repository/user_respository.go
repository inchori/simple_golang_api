package repository

import (
	"context"
	"grpc_identity/ent"
	"grpc_identity/ent/user"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, name, email, password string) (*ent.User, error)
	GetUserByID(ctx context.Context, id int) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	DeleteByID(ctx context.Context, id int) error
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
	return u.db.Get(ctx, id)
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.db.Query().Where(user.Email(email)).Only(ctx)
}

func (u *UserRepository) DeleteByID(ctx context.Context, id int) error {
	return u.db.DeleteOneID(id).Exec(ctx)
}
