package server

import (
	"context"
	"fmt"
	"grpc_identity/internal/middleware"
	"grpc_identity/internal/service"
	"grpc_identity/pb/v1beta1/auth"
	"grpc_identity/pkg/utils"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
		logrus.Errorf("failed to get user by email: %v", err)
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	if !utils.CheckPasswordHash(req.Password, userByEmail.Password) {
		logrus.Errorf("invalid password")
		return nil, fmt.Errorf("invalid password")
	}

	jwtToken, err := middleware.CreateAccessToken(strconv.Itoa(userByEmail.ID))
	if err != nil {
		logrus.Errorf("failed to create access jwt token: %v", err)
		return nil, fmt.Errorf("failed to create access jwt token: %v", err)
	}

	return &auth.LoginResponse{Token: jwtToken}, nil
}
