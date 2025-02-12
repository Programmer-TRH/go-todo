package main

import (
	"Todo-GO/config"
	"Todo-GO/handlers"
	"log"
	"net/http"
	"strings"
)

func main() {
	config.ConnectDB()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if path == "" || path == "/" {
			http.ServeFile(w, r, "static/index.html")
			return
		}

		filePath := "static/html/" + path + ".html"
		http.ServeFile(w, r, filePath)
	})

	http.HandleFunc("/todos", handlers.TodoHandler)
	http.HandleFunc("/add", handlers.AddTodoHandler)
	http.HandleFunc("/complete-todo", handlers.CompletedTodoHandler)
	http.HandleFunc("/delete-todo", handlers.DeleteTodoHandler)
	http.HandleFunc("/edit-todo", handlers.GetTodoHandler)
	http.HandleFunc("/update-todo", handlers.UpdateTodoHandler)

	log.Println("✅ Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
