package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Error handling with error types uses type assertions and type switches
// to check the specific type of an error and handle it accordingly

// Define different error types for different scenarios
type NetworkError struct {
	Code    int
	Message string
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error (code %d): %s", e.Code, e.Message)
}

type InputError struct {
	Field   string
	Message string
}

func (e InputError) Error() string {
	return fmt.Sprintf("invalid input for %s: %s", e.Field, e.Message)
}

type AuthError struct {
	Username string
	Message  string
}

func (e AuthError) Error() string {
	return fmt.Sprintf("authentication failed for %s: %s", e.Username, e.Message)
}

func main() {
	// Example 1: Handle different error types using type assertion
	err := fetchUserData("user123")
	if err != nil {
		// Type assertion to check for specific error type
		if netErr, ok := err.(NetworkError); ok {
			fmt.Printf("Network problem (code %d): %s\n", netErr.Code, netErr.Message)
			// We might retry the request based on the error code
			if netErr.Code == http.StatusServiceUnavailable {
				fmt.Println("Server is temporarily unavailable, retrying...")
			}
		} else {
			fmt.Println("Error fetching user data:", err)
		}
	}

	// Example 2: Handle different error types using type switch
	err = processUserLogin("alice", "password123")
	if err != nil {
		// Type switch to handle different error types
		switch e := err.(type) {
		case InputError:
			fmt.Printf("Please fix your input for field '%s': %s\n", e.Field, e.Message)
		case AuthError:
			fmt.Printf("Authentication problem for user '%s': %s\n", e.Username, e.Message)
		case NetworkError:
			fmt.Printf("Network issue (code %d): %s\n", e.Code, e.Message)
		default:
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Login successful!")
	}

	// Example 3: Using errors.As to check for error types (Go 1.13+)
	err = validateUserInput("bob", "")
	var inputErr InputError
	if errors.As(err, &inputErr) {
		fmt.Printf("Input validation failed: field '%s' has issue: %s\n", 
			inputErr.Field, inputErr.Message)
	}
}

// Function that simulates fetching user data
func fetchUserData(userID string) error {
	// Simulate a network error
	return NetworkError{
		Code:    http.StatusServiceUnavailable,
		Message: "server is overloaded",
	}
}

// Function that processes a user login
func processUserLogin(username, password string) error {
	if username == "" {
		return InputError{
			Field:   "username",
			Message: "cannot be empty",
		}
	}
	if password == "" {
		return InputError{
			Field:   "password",
			Message: "cannot be empty",
		}
	}
	if username == "alice" && password != "securepass" {
		return AuthError{
			Username: username,
			Message:  "invalid credentials",
		}
	}
	return nil
}

// Function that validates user input
func validateUserInput(username, password string) error {
	if username == "" {
		return InputError{
			Field:   "username",
			Message: "cannot be empty",
		}
	}
	if password == "" {
		return InputError{
			Field:   "password",
			Message: "cannot be empty",
		}
	}
	return nil
}
