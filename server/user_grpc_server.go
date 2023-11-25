package server

import (
	"context"
	"google.golang.org/grpc"
	"grpc_identity/pb/v1beta1/user"
	"grpc_identity/service"
)

type UserGRPCServiceServer struct {
	userService service.IUserService
	user.UnimplementedUserServer
}

func RegisterUserService(userService service.IUserService, svr *grpc.Server) {
	user.RegisterUserServer(svr, &UserGRPCServiceServer{userService: userService})
}

func (u *UserGRPCServiceServer) CreateGRPCUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	createUser, err := u.userService.CreateUser(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	userMsg := &user.UserMessage{
		Id:    int64(userByEmail.ID),
		Name:  userByEmail.Name,
		Email: userByEmail.Email,
	}

	return &user.GetUserByEmailResponse{User: userMsg}, nil
}

func (u *UserGRPCServiceServer) DeleteByID(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := u.userService.DeleteByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResponse{}, nil
}

func (u *UserGRPCServiceServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	updateUser, err := u.userService.UpdateUser(ctx, req.Name, req.Password, int(req.Id))
	if err != nil {
		return nil, err
	}

	userResp := &user.UpdateUserResponse{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	}

	return userResp, nil
}
