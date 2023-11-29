package server

import (
	"context"
	"fmt"
	"github.com/inchori/grpc_identity/pb/v1beta1/auth"
	"github.com/inchori/grpc_identity/pb/v1beta1/post"
	"github.com/inchori/grpc_identity/pb/v1beta1/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GatewayServer(grpcListenAddr string) (*runtime.ServeMux, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.DialContext(ctx, grpcListenAddr, opts...)
	if err != nil {
		logrus.Fatalf("failed to dial gRPC: %v", err)
		return nil, fmt.Errorf("failed to dial gRPC: %v", err)
	}

	err = auth.RegisterAuthHandler(ctx, mux, conn)
	if err != nil {
		logrus.Fatalf("failed to register auth handler gateway: %v", err)
		return nil, fmt.Errorf("failed to register auth handler gateway: %v", err)
	}

	err = user.RegisterUserHandler(ctx, mux, conn)
	if err != nil {
		logrus.Fatalf("failed to register user handler gateway: %v", err)
		return nil, fmt.Errorf("failed to register user handler gateway: %v", err)
	}

	err = post.RegisterPostHandler(ctx, mux, conn)
	if err != nil {
		logrus.Fatalf("failed to register post handler gateway: %v", err)
		return nil, fmt.Errorf("failed to register post handler gateway: %v", err)
	}

	return mux, nil
}
