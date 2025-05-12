package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Sentinel errors are predefined error values that can be directly compared
// They are useful for indicating specific error conditions that callers can check for

// Define our own sentinel errors
var (
	ErrInvalidInput   = errors.New("input is invalid")
	ErrNotFound       = errors.New("item not found")
	ErrNotAuthorized  = errors.New("not authorized")
)

func main() {
	// Example 1: Using a standard library sentinel error
	data, err := readData()
	if err == io.EOF {
		fmt.Println("Reached end of file")
	} else if err != nil {
		fmt.Println("Error reading data:", err)
	} else {
		fmt.Println("Data:", data)
	}

	// Example 2: Using our custom sentinel errors
	err = processItem(42)
	if err == ErrNotFound {
		fmt.Println("Item 42 was not found, please try another ID")
	} else if err == ErrNotAuthorized {
		fmt.Println("You are not authorized to access this item")
	} else if err != nil {
		fmt.Println("Unexpected error:", err)
	} else {
		fmt.Println("Item processed successfully")
	}

	// Example 3: Using errors.Is() to check for sentinel errors (Go 1.13+)
	err = validatePermissions("admin")
	if errors.Is(err, ErrNotAuthorized) {
		fmt.Println("Permission error:", err)
	} else if err != nil {
		fmt.Println("Other error:", err)
	} else {
		fmt.Println("Permissions valid")
	}
}

// Function that simulates reading data from a source
func readData() (string, error) {
	// Simulate EOF condition
	return "", io.EOF
}

// Function that processes an item by ID
func processItem(id int) error {
	// Simulate different error conditions based on ID
	if id < 0 {
		return ErrInvalidInput
	} else if id > 100 {
		return ErrNotFound
	} else if id == 42 {
		return ErrNotAuthorized
	}
	return nil
}

// Function that validates user permissions
func validatePermissions(role string) error {
	if role != "admin" {
		// We can wrap sentinel errors to add context while preserving the original error
		return fmt.Errorf("role '%s' has insufficient privileges: %w", role, ErrNotAuthorized)
	}
	return nil
}

// Function demonstrating how to check for file not found errors
func isFileNotFound(err error) bool {
	// Standard library provides some sentinel errors like os.ErrNotExist
	return errors.Is(err, os.ErrNotExist)
}
