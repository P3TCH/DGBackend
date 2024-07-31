package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/p3ch/DGBackend/cmd/api"
	"github.com/p3ch/DGBackend/config"
	"github.com/p3ch/DGBackend/db"
)

func main() {
	//database config
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	//connect to database
	initStorage(db)

	//create a new database connection
	server := api.NewAPIServer(":8080", db)

	//start the server with handler
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database: Connected")
}
