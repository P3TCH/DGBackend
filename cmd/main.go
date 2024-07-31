package main

import (
	"log"

	"github.com/p3ch/DGBackend/cmd/api"
)

func main() {
	//create a new database connection
	server := api.NewAPIServer(":8080", nil)

	//start the server with handler
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
