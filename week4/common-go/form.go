package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data. For application/x-www-form-urlencoded,
		// this populates r.Form and r.PostForm.
		// For multipart/form-data, use r.ParseMultipartForm(maxMemory int64).
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Access values using r.FormValue (convenience method that calls ParseForm)
		// It returns the first value for the specified key from r.Form.
		name := r.FormValue("name")
		email := r.FormValue("email")

		fmt.Fprintf(w, "Received POST request:\n")
		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Email: %s\n", email)

		// You can also access them directly from r.Form or r.PostForm
		fmt.Fprintf(w, "Raw Form map: %v\n", r.Form)
		fmt.Fprintf(w, "Raw PostForm map: %v\n", r.PostForm)

	} else {
		// Serve a simple HTML form for GET requests
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Form Example</title></head>
			<body>
				<form method="POST" action="/">
					<label for="name">Name:</label><br>
					<input type="text" id="name" name="name"><br>
					<label for="email">Email:</label><br>
					<input type="email" id="email" name="email"><br><br>
					<input type="submit" value="Submit">
				</form>
			</body>
			</html>
		`)
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}