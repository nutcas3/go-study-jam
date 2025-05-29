# Go Programming Examples

This repository contains practical examples demonstrating various aspects of Go programming from Week 4 of the Go course.

## Examples Overview

### 1. File Operations (`file_operations.go`)
Demonstrates basic file handling in Go:
- Writing data to files
- Reading file contents
- Getting file information (size, modification time)

```bash
go run file_operations.go
```

### 2. Networking (`networking.go`)
Shows TCP client/server communication:
- TCP server implementation
- Client connection handling
- Basic data transfer

```bash
go run networking.go
```

### 3. HTTP Server (`http_server.go`)
Implements a basic HTTP server with:
- Root endpoint with welcome message
- JSON API endpoint (/api/message)
- Support for GET and POST methods

```bash
# Start the server
go run http_server.go

# Test with curl
curl http://localhost:8080
curl http://localhost:8080/api/message
```

### 4. Concurrent Prime Number Finder
Find prime numbers using Go's concurrency features:
```bash
# Syntax:
go run cmd/main.go <start> <end>

# Example:
go run cmd/main.go 10 50
```

## Key Concepts Covered
- File I/O operations
- Network programming with TCP
- HTTP server implementation
- JSON handling
- Goroutines and channels
- Error handling
- Basic concurrency patterns

## Requirements
- Go 1.24 or higher
- No external dependencies required

## How to Run
Each example can be run independently using `go run`. Make sure you're in the correct directory when running the examples.
