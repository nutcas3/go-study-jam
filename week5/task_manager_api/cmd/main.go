package main

import (
	"log"
	"net/http"
	"strings"
	"task_manager_api/database"
	"task_manager_api/handlers"
)

func main() {
	// Initialize the database
	database.InitDB("tasks.db")
	
	// Set up the router
	http.HandleFunc("/tasks", tasksRouter)
	http.HandleFunc("/tasks/", tasksRouter)
	
	// Start the server
	log.Println("Task Manager API server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// tasksRouter routes all requests to the tasks handler
func tasksRouter(w http.ResponseWriter, r *http.Request) {
	// Simple routing based on the URL path
	if r.URL.Path == "/tasks" || strings.HasPrefix(r.URL.Path, "/tasks/") {
		handlers.TasksHandler(w, r)
		return
	}
	
	// If we get here, the path is not supported
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}
