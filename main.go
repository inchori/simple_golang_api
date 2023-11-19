package main

import (
	"context"
	"grpc_identity/config"
	"grpc_identity/database"
	"log"
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
}
