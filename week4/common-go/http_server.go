package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    // Handle root path
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to Go HTTP Server!")
    })

    // Handle JSON endpoint
    http.HandleFunc("/api/message", handleMessage)

    fmt.Println("Server starting on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        // Return a message
        message := Message{Text: "Hello from Go!"}
        json.NewEncoder(w).Encode(message)

    case "POST":
        // Read the incoming message
        var message Message
        if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Echo it back
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(message)

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
