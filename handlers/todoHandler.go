package handlers

import (
	"Todo-GO/config"
	"Todo-GO/models"
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var templates = template.Must(template.ParseFiles("templates/todos.html", "templates/edit.html"))

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var todos []models.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		log.Println("Error decoding todos:", err)
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d todos from database\n", len(todos))
	for _, todo := range todos {
		log.Printf("Todo: %+v\n", todo)
	}
	if len(todos) == 0 {
		log.Println("⚠️ No todos found in the database.")
		http.Error(w, "No todos found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = templates.ExecuteTemplate(w, "todos", todos)
	if err != nil {
		log.Println("❌ Template rendering error:", err)
		http.Error(w, "Error rendering HTML", http.StatusInternalServerError)
	}
}
