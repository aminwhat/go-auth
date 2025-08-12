package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo(uri string, username string, password string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoOptions := options.Client().ApplyURI(uri)
	mongoOptions.SetAuth(options.Credential{
		Username:      username,
		Password:      password,
		AuthSource:    "admin",
		AuthMechanism: "SCRAM-SHA-256",
	})

	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB")
	return client
}
