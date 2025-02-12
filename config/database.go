package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func ConnectDB() {
	clinetOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clinetOptions)
	if err != nil {
		log.Fatal("❌ MongoDB Connection Failed:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("❌ MongoDB Ping Failed:", err)
	}

	DB = client.Database("Todo-GO").Collection("todos")
	log.Println("✅ Connected to MongoDB!")
}
