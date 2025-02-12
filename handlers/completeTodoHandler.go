package handlers

import (
	"Todo-GO/config"
	"Todo-GO/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CompletedTodoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "❌ Invalid request method!", http.StatusMethodNotAllowed)
		return
	}
	todoID := r.FormValue("id")
	fmt.Println("🔍 Received ID:", todoID) // Debugging
	objID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		http.Error(w, "❌ Invalid ID!", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.DB.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"completed": true}})
	if err != nil {
		http.Error(w, "❌ Invalid ID!", http.StatusBadRequest)
		return
	}

	log.Println("✅Todo marked as completed:", todoID)

	cursor, err := config.DB.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "❌ Database Error", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	var todos []models.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		http.Error(w, "❌ Error fetching todos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = templates.ExecuteTemplate(w, "todos", todos)
	if err != nil {
		http.Error(w, "❌ Error rendering todos", http.StatusInternalServerError)
	}

}
