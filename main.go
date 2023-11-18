package main

import (
	"grpc_identity/config"
	"grpc_identity/pkg/database"
	"log"
)

func main() {
	config, err := config.LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	database.ConnectDB(config)
}
