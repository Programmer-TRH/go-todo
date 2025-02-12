package config

import (
	"context"
	"fmt"
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

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("⚠️ MongoDB Disconnect Failed:", err)
		}
	}()

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("❌ MongoDB Ping Failed:", err)
	}

	DB = client.Database("Todo-GO").Collection("todos")
	fmt.Println("✅ Connected to MongoDB!")
}
