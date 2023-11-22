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

	middleware.NewLoginHandler(app.Group("/v1/auth"))

	authentication := middleware.NewAuthentication()

	handler.NewUserHandler(app.Group("/v1/users"), context.Background(), userService, authentication.Authentication())
	handler.NewPostHandler(app.Group("/v1/posts"), context.Background(), postService, authentication.Authentication())

	log.Fatal(app.Listen(":3000"))
}
