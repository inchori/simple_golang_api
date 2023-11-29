package database

import (
	"fmt"
	"github.com/inchori/grpc_identity/config"
	"github.com/inchori/grpc_identity/internal/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *ent.Client

func ConnectDB(config config.Config) *ent.Client {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DBConn = client
	return DBConn
}
