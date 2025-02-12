package handlers

import (
	"Todo-GO/config"
	"Todo-GO/models"
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "❌ Invalid request method!", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")

	if title == "" || description == "" {
		http.Error(w, "❌ Title and Description required!", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newTodo := models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
	}

	_, err := config.DB.InsertOne(ctx, newTodo)
	if err != nil {
		http.Error(w, "❌ Failed to save to database!", http.StatusInternalServerError)
		return
	}

	log.Println("✅ To-Do Added:", title)

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
	templates.ExecuteTemplate(w, "todos", todos)

}
