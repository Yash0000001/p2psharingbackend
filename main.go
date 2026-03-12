package main

import (
	"log"
	"net/http"

	"github.com/yash0000001/p2psharingbackend/internal/database"
)

func main() {
	database.DBConnect()

	log.Println("Server started on port 8080")

	http.ListenAndServe(":8080", nil)
}
