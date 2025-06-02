package main

import (
	"crud_api/database"
	"crud_api/handlers"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Initialize the database
	database.InitDB("items.db")
	
	// Set up the router
	http.HandleFunc("/items", itemsRouter)
	http.HandleFunc("/items/", itemsRouter)
	
	// Start the server
	log.Println("CRUD API server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// itemsRouter routes all requests to the items handler
func itemsRouter(w http.ResponseWriter, r *http.Request) {
	// Simple routing based on the URL path
	if r.URL.Path == "/items" || strings.HasPrefix(r.URL.Path, "/items/") {
		handlers.ItemsHandler(w, r)
		return
	}
	
	// If we get here, the path is not supported
	w.WriteHeader(http.StatusNotFound)
}
