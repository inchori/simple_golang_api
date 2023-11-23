package main

import (
	"context"
	"grpc_identity/config"
	"grpc_identity/database"
	"grpc_identity/handler"
	"grpc_identity/middleware"
	"grpc_identity/repository"
	"grpc_identity/service"
	"log"

	"github.com/gofiber/fiber/v2"
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

	protected := middleware.Protected()

	handler.NewLoginHandler(app.Group("/v1/auth"), context.Background(), userService)
	handler.NewUserHandler(app.Group("/v1/users"), context.Background(), userService, protected)
	handler.NewPostHandler(app.Group("/v1/posts"), context.Background(), postService, userService, protected)

	log.Fatal(app.Listen(":3000"))
}
