package server

import (
	"context"
	"google.golang.org/grpc"
	"grpc_identity/middleware"
	"grpc_identity/pb/v1beta1/auth"
	"grpc_identity/service"
	"grpc_identity/utils"
	"strconv"
)

type AuthGRPCServiceServer struct {
	userService service.IUserService
	auth.UnimplementedAuthServer
}

func RegisterAuthService(userService service.IUserService, svr *grpc.Server) {
	auth.RegisterAuthServer(svr, &AuthGRPCServiceServer{
		userService: userService,
	})
}

func (a *AuthGRPCServiceServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	userByEmail, err := a.userService.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, userByEmail.Password) {
		return nil, err
	}

	jwtToken, err := middleware.CreateAccessToken(strconv.Itoa(userByEmail.ID))
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{Token: jwtToken}, nil
}
