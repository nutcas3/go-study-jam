package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// This example demonstrates a practical application that combines multiple
// error handling techniques in a file processing application

// Define sentinel errors
var (
	ErrFileNotFound    = errors.New("file not found")
	ErrInvalidFormat   = errors.New("invalid file format")
	ErrProcessingLimit = errors.New("processing limit exceeded")
)

// Define custom error types with additional context
type FileError struct {
	Filename string
	Op       string
	Err      error
}

func (e FileError) Error() string {
	return fmt.Sprintf("file %s: %s error: %v", e.Filename, e.Op, e.Err)
}

// Implement Unwrap to support error wrapping
func (e FileError) Unwrap() error {
	return e.Err
}

// Define an error with behavior
type RetryableError struct {
	Err       error
	MaxTries  int
	Delay     time.Duration
	Temporary bool
}

func (e RetryableError) Error() string {
	return fmt.Sprintf("%v (retryable: max tries=%d, delay=%v)", e.Err, e.MaxTries, e.Delay)
}

func (e RetryableError) Unwrap() error {
	return e.Err
}

func (e RetryableError) IsTemporary() bool {
	return e.Temporary
}

func (e RetryableError) ShouldRetry() bool {
	return true
}

func (e RetryableError) RetryAfter() time.Duration {
	return e.Delay
}

// Main function demonstrating combined error handling approaches
func main() {
	fmt.Println("File Processing Application")
	fmt.Println("==========================")

	// Process a file with combined error handling
	err := processFile("data.txt")
	if err != nil {
		// 1. Check for sentinel errors using errors.Is
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("Error: The specified file could not be found.")
			fmt.Println("Please check the filename and try again.")
		} else if errors.Is(err, ErrInvalidFormat) {
			fmt.Println("Error: The file format is invalid.")
			fmt.Println("Please ensure the file is a valid text file.")
		} else if errors.Is(err, ErrProcessingLimit) {
			fmt.Println("Error: Processing limit exceeded.")
			fmt.Println("Try processing a smaller file or increase the limit.")
		} else {
			// 2. Check for specific error types using errors.As
			var fileErr FileError
			if errors.As(err, &fileErr) {
				fmt.Printf("File operation error: %s operation on %s failed\n", 
					fileErr.Op, fileErr.Filename)
			}

			// 3. Check for retryable errors
			var retryErr RetryableError
			if errors.As(err, &retryErr) {
				if retryErr.IsTemporary() {
					fmt.Printf("This is a temporary error that can be retried.\n")
					fmt.Printf("Recommended: retry up to %d times with %v delay between attempts\n", 
						retryErr.MaxTries, retryErr.RetryAfter())
					
					// Simulate retry logic
					simulateRetry(retryErr.MaxTries, retryErr.RetryAfter(), "processFile")
				}
			}

			// 4. Print the full error chain
			fmt.Println("\nError details:")
			printErrorChain(err)
		}
	} else {
		fmt.Println("File processed successfully!")
	}
}

// Function to process a file with various error scenarios
func processFile(filename string) error {
	// Step 1: Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Wrap the sentinel error with context
		return FileError{
			Filename: filename,
			Op:       "stat",
			Err:      ErrFileNotFound,
		}
	}

	// Step 2: Open the file (this will fail since we're using a non-existent file)
	file, err := os.Open(filename)
	if err != nil {
		// Make this a retryable error
		return RetryableError{
			Err:       fmt.Errorf("opening file: %w", err),
			MaxTries:  3,
			Delay:     2 * time.Second,
			Temporary: true,
		}
	}
	defer file.Close()

	// Step 3: Read and validate file format
	valid, err := validateFileFormat(file)
	if err != nil {
		return fmt.Errorf("validating file format: %w", err)
	}
	if !valid {
		return FileError{
			Filename: filename,
			Op:       "validate",
			Err:      ErrInvalidFormat,
		}
	}

	// Step 4: Process file content
	err = processFileContent(file)
	if err != nil {
		return fmt.Errorf("processing file content: %w", err)
	}

	return nil
}

// Function to validate file format
func validateFileFormat(file *os.File) (bool, error) {
	// This is a simplified example
	// In a real application, we would check file headers, etc.
	return true, nil
}

// Function to process file content
func processFileContent(file *os.File) error {
	// Simplified example
	_, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	
	// Simulate a processing limit error
	return ErrProcessingLimit
}

// Helper function to print the full error chain
func printErrorChain(err error) {
	if err == nil {
		return
	}
	
	// Print the current error
	fmt.Printf("- %v\n", err)
	
	// Recursively print the wrapped error
	printErrorChain(errors.Unwrap(err))
}

// Helper function to simulate retry logic
func simulateRetry(maxTries int, delay time.Duration, operation string) {
	for i := 1; i <= maxTries; i++ {
		fmt.Printf("Retry %d/%d: Attempting %s again...\n", i, maxTries, operation)
		// In a real application, we would actually retry the operation here
		
		// Simulate a successful retry on the last attempt
		if i == maxTries {
			fmt.Println("Retry successful!")
			return
		} else {
			fmt.Printf("Retry failed, waiting %v before next attempt\n", delay)
			// In a real application, we would actually wait here
		}
	}
	fmt.Println("All retry attempts failed")
}
