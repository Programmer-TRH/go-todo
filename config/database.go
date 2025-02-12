package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func ConnectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clinetOptions := options.Client().ApplyURI("mongodb+srv://tokdirrahoman4:01833740078Trh@todo-go.qksdh.mongodb.net/?retryWrites=true&w=majority&appName=Todo-GO").SetServerAPIOptions(serverAPI)
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
