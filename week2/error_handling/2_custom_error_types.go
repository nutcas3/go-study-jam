package main

import (
	"fmt"
)

// Custom error types provide more context and information about errors
// They allow for more detailed error handling and reporting

// Define a custom error type for validation errors
type ValidationError struct {
	Field string
	Issue string
}

// Implement the error interface for our custom error type
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field %s: %s", e.Field, e.Issue)
}

// Define another custom error type for database errors
type DBError struct {
	Operation string
	Message   string
}

// Implement the error interface for our database error type
func (e DBError) Error() string {
	return fmt.Sprintf("database %s error: %s", e.Operation, e.Message)
}

func main() {
	// Example 1: Validate user input
	err := validateUser("", "password123")
	if err != nil {
		fmt.Println("User validation error:", err)
	}

	// Example 2: Simulate a database operation
	user, err := getUserFromDB(123)
	if err != nil {
		fmt.Println("Database error:", err)
		
		// We can also check the specific type of error
		if dbErr, ok := err.(DBError); ok {
			fmt.Printf("DB Operation that failed: %s\n", dbErr.Operation)
		}
	} else {
		fmt.Println("User found:", user)
	}
}

// Function that validates user input and returns a custom error
func validateUser(username, password string) error {
	if username == "" {
		return ValidationError{
			Field: "username",
			Issue: "cannot be empty",
		}
	}
	if len(password) < 8 {
		return ValidationError{
			Field: "password",
			Issue: "must be at least 8 characters",
		}
	}
	return nil
}

// Function that simulates fetching a user from a database
func getUserFromDB(id int) (string, error) {
	// Simulate a database error
	if id > 100 {
		return "", DBError{
			Operation: "query",
			Message:   "user not found",
		}
	}
	return fmt.Sprintf("User-%d", id), nil
}
