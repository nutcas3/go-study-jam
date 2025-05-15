package main

import (
	"fmt"
	"time"
)

// Error handling with behavior uses interfaces to define error behaviors
// This approach focuses on what the error can do rather than what it is

// Define interfaces for different error behaviors
type Temporary interface {
	Temporary() bool
}

type Timeout interface {
	Timeout() bool
}

type Retryable interface {
	Retryable() bool
	MaxRetries() int
}

// Define error types that implement these behaviors
type ServiceError struct {
	Message      string
	IsTemporary  bool
	IsTimeout    bool
	CanRetry     bool
	RetryCount   int
	OccurredAt   time.Time
}

// Implement the error interface
func (e ServiceError) Error() string {
	return fmt.Sprintf("%s (occurred at: %s)", e.Message, e.OccurredAt.Format(time.RFC3339))
}

// Implement behavior interfaces
func (e ServiceError) Temporary() bool {
	return e.IsTemporary
}

func (e ServiceError) Timeout() bool {
	return e.IsTimeout
}

func (e ServiceError) Retryable() bool {
	return e.CanRetry
}

func (e ServiceError) MaxRetries() int {
	return e.RetryCount
}

func main() {
	// Example 1: Handle a temporary service error
	err := connectToService("api.example.com")
	if err != nil {
		fmt.Println("Service connection error:", err)
		
		// Check if the error is temporary
		if temp, ok := err.(Temporary); ok && temp.Temporary() {
			fmt.Println("This is a temporary error, we can retry later")
		}
		
		// Check if the error is a timeout
		if timeout, ok := err.(Timeout); ok && timeout.Timeout() {
			fmt.Println("Connection timed out, we might retry with a longer timeout")
		}
		
		// Check if the error is retryable
		if retry, ok := err.(Retryable); ok && retry.Retryable() {
			maxRetries := retry.MaxRetries()
			fmt.Printf("This error is retryable, we can retry up to %d times\n", maxRetries)
		}
	}

	// Example 2: Process a database operation with potential errors
	err = performDatabaseOperation("INSERT")
	handleOperationError(err)
}

// Function that simulates connecting to a service
func connectToService(url string) error {
	// Simulate a temporary network timeout
	return ServiceError{
		Message:     "failed to connect to service",
		IsTemporary: true,
		IsTimeout:   true,
		CanRetry:    true,
		RetryCount:  3,
		OccurredAt:  time.Now(),
	}
}

// Function that simulates a database operation
func performDatabaseOperation(operation string) error {
	// Simulate different errors based on the operation
	if operation == "INSERT" {
		return ServiceError{
			Message:     "database connection lost during insert",
			IsTemporary: true,
			IsTimeout:   false,
			CanRetry:    true,
			RetryCount:  5,
			OccurredAt:  time.Now(),
		}
	}
	return nil
}

// Function that handles operation errors based on their behavior
func handleOperationError(err error) {
	if err == nil {
		fmt.Println("Operation completed successfully")
		return
	}
	
	fmt.Println("Operation failed:", err)
	
	// Implement a retry strategy based on error behavior
	var retryDelay time.Duration = 1 * time.Second
	var maxRetries int = 1
	
	if temp, ok := err.(Temporary); ok && temp.Temporary() {
		fmt.Println("Error is temporary, implementing retry strategy")
		
		if retry, ok := err.(Retryable); ok && retry.Retryable() {
			maxRetries = retry.MaxRetries()
			fmt.Printf("Will retry up to %d times\n", maxRetries)
			
			// Simulate retry logic
			for i := 1; i <= maxRetries; i++ {
				fmt.Printf("Retry attempt %d after %v delay...\n", i, retryDelay)
				// In a real application, we would wait and retry the operation
				retryDelay *= 2 // Exponential backoff
			}
		}
	} else {
		fmt.Println("Error is permanent, no retry possible")
	}
}
