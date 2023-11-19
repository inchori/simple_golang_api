package main

import (
	"context"
	"grpc_identity/config"
	"grpc_identity/database"
	"grpc_identity/handler"
	"grpc_identity/repository"
	"grpc_identity/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	loadConfig, err := config.LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.ConnectDB(loadConfig)
	ctx := context.Background()
	if err := dbClient.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	userRepository := repository.NewUserRepository(dbClient.User)
	userService := service.NewUserService(userRepository)
	handler.NewUserHandler(app.Group("/v1/users"), context.Background(), userService)

	log.Fatal(app.Listen(":3000"))
}
