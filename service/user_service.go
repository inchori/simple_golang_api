package service

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(ctx context.Context, name, email, password string) (dto.UserResponse, error)
	GetUserByID(ctx context.Context, id int) (dto.UserResponse, error)
	GetUserByName(ctx context.Context, name string) (dto.UserResponse, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateByName(ctx context.Context, name string, id int) (dto.UserResponse, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{repo: userRepository}
}

func (u *UserService) CreateUser(ctx context.Context, name, email, password string) (dto.UserResponse, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return dto.UserResponse{}, err
	}
	user, err := u.repo.CreateUser(ctx, name, email, string(encryptedPassword))
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

func (u *UserService) DeleteByID(ctx context.Context, id int) error {
	return u.repo.DeleteByID(ctx, id)
}

func (u *UserService) UpdateByName(ctx context.Context, name string, id int) (dto.UserResponse, error) {
	updateUserByName, err := u.repo.UpdateUserByName(ctx, name, id)
	if err != nil {
		return dto.UserResponse{}, nil
	}

	userResponse := dto.NewUserResponse(updateUserByName)
	return userResponse, nil
}
