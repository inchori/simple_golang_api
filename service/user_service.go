package service

import (
	"context"
	"grpc_identity/ent"
	"grpc_identity/repository"
)

type IUserService interface {
	CreateUser(ctx context.Context, name, email, password string) (*ent.User, error)
	GetUserByID(ctx context.Context, id int) (*ent.User, error)
	GetUserByName(ctx context.Context, name string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, name, password string, id int) (*ent.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{repo: userRepository}
}

func (u *UserService) CreateUser(ctx context.Context, name, email, password string) (*ent.User, error) {
	return u.repo.CreateUser(ctx, name, email, password)
}

func (u *UserService) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *UserService) GetUserByName(ctx context.Context, name string) (*ent.User, error) {
	return u.repo.GetUserByName(ctx, name)
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *UserService) DeleteByID(ctx context.Context, id int) error {
	return u.repo.DeleteByID(ctx, id)
}

func (u *UserService) UpdateUser(ctx context.Context, name, password string, id int) (*ent.User, error) {
	return u.repo.UpdateUser(ctx, name, password, id)
}
