package server

import (
	"context"
	"fmt"
	"grpc_identity/internal/service"
	"grpc_identity/pb/v1beta1/user"
	"grpc_identity/pkg/utils"
	"grpc_identity/server/interceptor"
	"strconv"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type UserGRPCServiceServer struct {
	userService service.IUserService
	user.UnimplementedUserServer
}

func RegisterUserService(userService service.IUserService, svr *grpc.Server) {
	user.RegisterUserServer(svr, &UserGRPCServiceServer{
		userService: userService,
	})
}

func (u *UserGRPCServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logrus.Errorf("failed to hash password: %v", err)
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	createUser, err := u.userService.CreateUser(ctx, req.Name, req.Email, hashPassword)
	if err != nil {
		logrus.Errorf("failed to create user: %v", err)
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	userResp := &user.CreateUserResponse{
		Id:    int64(createUser.ID),
		Email: createUser.Email,
		Name:  createUser.Name,
	}

	return userResp, nil
}

func (u *UserGRPCServiceServer) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error) {
	userByID, err := u.userService.GetUserByID(ctx, int(req.Id))
	if err != nil {
		logrus.Errorf("failed to get user by ID: %v", err)
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	userMsg := &user.UserMessage{
		Id:    int64(userByID.ID),
		Name:  userByID.Name,
		Email: userByID.Email,
	}

	return &user.GetUserByIDResponse{User: userMsg}, nil
}

func (u *UserGRPCServiceServer) GetUserByName(ctx context.Context, req *user.GetUserByNameRequest) (*user.GetUserByNameResponse, error) {
	userByName, err := u.userService.GetUserByName(ctx, req.Name)
	if err != nil {
		logrus.Errorf("failed to get user by name: %v", err)
		return nil, fmt.Errorf("failed to get user by name: %v", err)
	}

	userMsg := &user.UserMessage{
		Id:    int64(userByName.ID),
		Name:  userByName.Name,
		Email: userByName.Email,
	}

	return &user.GetUserByNameResponse{User: userMsg}, nil
}

func (u *UserGRPCServiceServer) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {
	userByEmail, err := u.userService.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logrus.Errorf("failed to get user by email: %v", err)
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	userMsg := &user.UserMessage{
		Id:    int64(userByEmail.ID),
		Name:  userByEmail.Name,
		Email: userByEmail.Email,
	}

	return &user.GetUserByEmailResponse{User: userMsg}, nil
}

func (u *UserGRPCServiceServer) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	jwtToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		logrus.Errorf("failed to get jwt token: %v", err)
		return nil, fmt.Errorf("failed to get jwt token: %v", err)
	}

	tokenClaimsID, err := interceptor.ExtractTokenFromMetadata(jwtToken)
	if err != nil {
		logrus.Errorf("failed to extract token from metadata: %v", err)
		return nil, fmt.Errorf("failed to extract token from metadata: %v", err)
	}

	tokenID, _ := strconv.Atoi(tokenClaimsID)
	userByID, err := u.userService.GetUserByID(ctx, tokenID)
	if err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if userByID.ID != tokenID {
		logrus.Errorf("unauthorized")
		return nil, fmt.Errorf("unauthorzied")
	}

	err = u.userService.DeleteByID(ctx, int(req.Id))
	if err != nil {
		logrus.Errorf("failed to delete user: %v", err)
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}

	return &user.DeleteUserResponse{}, nil
}

func (u *UserGRPCServiceServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	jwtToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		logrus.Errorf("failed to get jwt token: %v", err)
		return nil, fmt.Errorf("failed to get jwt token: %v", err)
	}

	tokenClaimsID, err := interceptor.ExtractTokenFromMetadata(jwtToken)
	if err != nil {
		logrus.Errorf("failed to extract token from metadata: %v", err)
		return nil, fmt.Errorf("failed to extract token from metadata: %v", err)
	}

	tokenID, _ := strconv.Atoi(tokenClaimsID)
	userByID, err := u.userService.GetUserByID(ctx, tokenID)
	if err != nil {
		logrus.Errorf("failed to get user by ID: %v", err)
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	if userByID.ID != tokenID {
		logrus.Errorf("unauthorized")
		return nil, fmt.Errorf("unauthorzied")
	}
	updateUser, err := u.userService.UpdateUser(ctx, req.Name, req.Password, int(req.Id))
	if err != nil {
		logrus.Errorf("failed to update user: %v", err)
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	userResp := &user.UpdateUserResponse{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	}

	return userResp, nil
}
