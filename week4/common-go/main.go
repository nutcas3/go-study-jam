package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Query() returns a url.Values, which is a map[string][]string
	queryParams := r.URL.Query()

	// Get a single value for a parameter
	query := queryParams.Get("query") // Returns the first value for "query"
	if query != "" {
		fmt.Fprintf(w, "Query: %s\n", query)
	} else {
		fmt.Fprintf(w, "No query parameter found.\n")
	}

	// Get all values for a parameter (if it can appear multiple times)
	pages := queryParams["page"] // Returns a slice of strings
	if len(pages) > 0 {
		fmt.Fprintf(w, "Pages: %v\n", pages)
	}

	// Check if a parameter exists
	if _, ok := queryParams["debug"]; ok {
		fmt.Fprintf(w, "Debug mode is enabled.\n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}