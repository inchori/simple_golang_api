package service

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/repository"
)

type IUserService interface {
	CreateUser(ctx context.Context, name, email, password string) (dto.UserResponse, error)
	GetUserByID(ctx context.Context, id int) (dto.UserResponse, error)
	GetUserByName(ctx context.Context, name string) (dto.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, string, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, name, password string, id int) (dto.UserResponse, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{repo: userRepository}
}

func (u *UserService) CreateUser(ctx context.Context, name, email, password string) (dto.UserResponse, error) {
	user, err := u.repo.CreateUser(ctx, name, email, password)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userResponse := dto.NewUserResponse(user)
	return userResponse, nil
}

func (u *UserService) GetUserByID(ctx context.Context, id int) (dto.UserResponse, error) {
	userByID, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	userResponse := dto.NewUserResponse(userByID)
	return userResponse, nil
}

func (u *UserService) GetUserByName(ctx context.Context, name string) (dto.UserResponse, error) {
	userByName, err := u.repo.GetUserByName(ctx, name)
	if err != nil {
		return dto.UserResponse{}, err
	}

	userResponse := dto.NewUserResponse(userByName)
	return userResponse, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, string, error) {
	userByEmail, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.UserResponse{}, "", err
	}

	userResponse := dto.NewUserResponse(userByEmail)
	return userResponse, userByEmail.Password, nil
}

func (u *UserService) DeleteByID(ctx context.Context, id int) error {
	return u.repo.DeleteByID(ctx, id)
}

func (u *UserService) UpdateUser(ctx context.Context, name, password string, id int) (dto.UserResponse, error) {
	updateUserByName, err := u.repo.UpdateUser(ctx, name, password, id)
	if err != nil {
		return dto.UserResponse{}, nil
	}

	userResponse := dto.NewUserResponse(updateUserByName)
	return userResponse, nil
}
