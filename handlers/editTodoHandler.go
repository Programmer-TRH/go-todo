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

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("🔍 Received ID:", id) // Debugging
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "❌ Invalid ID!", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todo models.Todo
	err = config.DB.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		http.Error(w, "❌ Task not found!", http.StatusNotFound)
		return
	}
	templates.ExecuteTemplate(w, "edit-form", todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "❌ Invalid request method!", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("description")

	fmt.Println("🔍 Received ID:", id) // Debugging

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "❌ Invalid ID!", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.DB.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"title": title, "description": description}})
	if err != nil {
		http.Error(w, "❌ Failed to update!", http.StatusInternalServerError)
		return
	}

	log.Println("✅Todo Updated:", title)

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
