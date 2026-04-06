package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func DBConnect() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

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

	DB = client.Database("p2psharing")
}

func CreateUserIndexes() {

	collection := DB.Collection("users")

	indexes := []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := collection.Indexes().CreateMany(context.Background(), indexes)

	if err != nil {
		log.Println("Index creation error:", err)
	}
}
