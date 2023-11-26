package main

import (
	"context"
	"grpc_identity/config"
	"grpc_identity/database"
	"grpc_identity/handler"
	"grpc_identity/middleware"
	"grpc_identity/repository"
	"grpc_identity/server"
	"grpc_identity/server/interceptor"
	"grpc_identity/service"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	app := fiber.New()

	loadConfig, err := config.LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.ConnectDB(loadConfig)
	ctx := context.Background()
	if err := dbClient.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(dbClient.User)
	userService := service.NewUserService(userRepository)
	postRepository := repository.NewPostRepository(dbClient.Post)
	postService := service.NewPostService(postRepository)

	if loadConfig.Server == "grpc" {
		lis, err := net.Listen("tcp", ":"+loadConfig.GRPCPort)
		if err != nil {
			log.Fatal(err)
		}

		methods := []string{
			"proto.v1beta1.post.Post/UpdatePost",
			"proto.v1beta1.post.Post/DeletePost",
			"proto.v1beta1.user.Post/GetPostByUser",
			"proto.v1beta1.user.User/UpdateUser",
			"proto.v1beta1.user.User/DeleteUser",
		}

		logger := logrus.New()

		jwtInterceptor := interceptor.NewJWTInterceptor(methods)

		grpcSvr := grpc.NewServer(grpc.ChainUnaryInterceptor(
			jwtInterceptor.Interceptor,
			logging.UnaryServerInterceptor(interceptor.LoggerInterceptor(logger)),
		))

		server.RegisterAuthService(userService, grpcSvr)
		server.RegisterUserService(userService, grpcSvr)
		server.RegisterPostService(postService, userService, grpcSvr)
		reflection.Register(grpcSvr)

		logger.Infof("gRPC server is running on %s port", loadConfig.GRPCPort)
		if err := grpcSvr.Serve(lis); err != nil {
			log.Fatal(err)
		}

	} else if loadConfig.Server == "http" {
		protected := middleware.Protected()

		handler.NewLoginHandler(app.Group("/v1/auth"), context.Background(), userService)
		handler.NewUserHandler(app.Group("/v1/users"), context.Background(), userService, protected)
		handler.NewPostHandler(app.Group("/v1/posts"), context.Background(), postService, userService, protected)
		log.Fatal(app.Listen(":" + loadConfig.HTTPPort))
	}

}
