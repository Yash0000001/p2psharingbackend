package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func DBConnect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Database Connection Error:", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Database Ping Failed:", err)
	}

	log.Println("Database Connected Successfully nigga!!")

	DB = client.Database("p2p sharing")
}
