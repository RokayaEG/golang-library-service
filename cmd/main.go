package main

import (
	"fmt"
	"log"

	"github.com/RokayaEG/golang-library-service/cmd/api"
	"github.com/RokayaEG/golang-library-service/config"
	database "github.com/RokayaEG/golang-library-service/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPasswd,
		Addr:   fmt.Sprintf("%s:%s", config.Envs.DBHost, config.Envs.DBPort),
		DBName: config.Envs.DBName,
	})

	database.InitStorage(db)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(config.Envs.Port, db)

	server.Run()

}
