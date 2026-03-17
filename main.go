package main

import (
	"log"
	"net/http"

	"github.com/yash0000001/p2psharingbackend/internal/database"
	"github.com/yash0000001/p2psharingbackend/internal/routes"
)

func main() {
	database.DBConnect()
	database.CreateUserIndexes()
	log.Println("Server started on port 8080")


	// routes 
	routes.AuthRoutes()

	http.ListenAndServe(":8080", nil)
}
