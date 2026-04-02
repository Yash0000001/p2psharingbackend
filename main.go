package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yash0000001/p2psharingbackend/internal/database"
	"github.com/yash0000001/p2psharingbackend/internal/routes"
	"github.com/yash0000001/p2psharingbackend/internal/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	database.DBConnect()
	database.CreateUserIndexes()
	log.Println("Server started on port 8080")

	// routes
	routes.AuthRoutes()
	routes.RoomRoutes()
	routes.SignallingRoutes()

	http.ListenAndServe(":8080", utils.EnableCORS(http.DefaultServeMux))
}
