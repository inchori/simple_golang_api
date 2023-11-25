package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_identity/config"
	"grpc_identity/database"
	"grpc_identity/handler"
	"grpc_identity/middleware"
	"grpc_identity/repository"
	"grpc_identity/server"
	"grpc_identity/service"
	"log"
	"net"
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
		//conn, err := grpc.DialContext(context.Background(), "localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
		//if err != nil {
		//	log.Fatal(err)
		//}
		lis, err := net.Listen("tcp", ":4040")
		if err != nil {
			log.Fatal(err)
		}

		grpcSvr := grpc.NewServer()
		server.RegisterUserService(userService, grpcSvr)
		server.RegisterPostService(postService, userService, grpcSvr)
		reflection.Register(grpcSvr)

		if err := grpcSvr.Serve(lis); err != nil {
			log.Fatal(err)
		}

	} else if loadConfig.Server == "http" {
		protected := middleware.Protected()

		handler.NewLoginHandler(app.Group("/v1/auth"), context.Background(), userService)
		handler.NewUserHandler(app.Group("/v1/users"), context.Background(), userService, protected)
		handler.NewPostHandler(app.Group("/v1/posts"), context.Background(), postService, userService, protected)
		log.Fatal(app.Listen(":3000"))
	}

	//srv := grpc.NewServer()

	//userGRPCHandler := server.NewUserGRPCHandler(userService)
	//user.RegisterUserServer(srv, &userGRPCHandler)

	//if err := srv.Serve(lis); err != nil {
	//	log.Fatal(err)
	//}

}
