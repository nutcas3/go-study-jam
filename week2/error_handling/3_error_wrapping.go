package main

import (
	"errors"
	"fmt"
	"os"
)

// Error wrapping in Go allows you to add context to errors while preserving the original error
// This helps in creating error chains that provide more information about what went wrong

func main() {
	// Example: Try to process a configuration file
	err := processConfig("config.json")
	if err != nil {
		fmt.Println("Configuration error:", err)
		
		// We can unwrap the error to get the original cause
		fmt.Println("\nError chain:")
		for err != nil {
			fmt.Printf("- %v\n", err)
			err = errors.Unwrap(err)
		}
	}
}


// Function that processes a configuration file
func processConfig(filename string) error {
	// Try to read the file
	data, err := readConfigFile(filename)
	if err != nil {
		// Wrap the error with additional context using %w verb
		return fmt.Errorf("failed to process config: %w", err)
	}
	
	// Try to parse the data
	err = parseConfigData(data)
	if err != nil {
		return fmt.Errorf("config processing error: %w", err)
	}
	
	return nil
}

// Function that reads a configuration file
func readConfigFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		// Wrap the error with additional context
		return nil, fmt.Errorf("error reading config file %s: %w", filename, err)
	}
	return data, nil
}

// Function that parses configuration data
func parseConfigData(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty configuration data")
	}
	
	// Simulate a parsing error
	return fmt.Errorf("invalid JSON format at position %d", 42)
}

// Helper function to demonstrate how to check for specific wrapped errors
func isFileNotFoundError(err error) bool {
	// errors.Is traverses the error chain created by fmt.Errorf with %w
	return errors.Is(err, os.ErrNotExist)
}
