package main

import (
	"errors"
	"fmt"
	"os"
)

// Basic error handling in Go using the standard error interface
// This is the most common approach in Go

func main() {
	// Example 1: Simple error checking
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result of 10/2:", result)

	// Example 2: Error with zero division
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		// We continue execution instead of returning
	}


	// Example 3: Reading a file that doesn't exist
	content, err := readFile("non_existent_file.txt")
	if err != nil {
		fmt.Println("File error:", err)
		// We can log the error and continue or exit
	} else {
		fmt.Println("File content:", content)
	}

	fmt.Println("Program completed.")
}

// Function that returns an error when dividing by zero
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Function that reads a file and returns its content or an error
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err // Return the error as is
	}
	return string(data), nil
}
