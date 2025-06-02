# Task Manager REST API

This project is part of Week 5 of the Go course. It is a REST API for a Task Manager application with unit tests, demonstrating:
- Writing unit tests
- Table-driven tests
- Code organization and best practices
- SQL database integration

## Project Structure

```
task_manager_api/
├── cmd/
│   └── main.go           # Application entry point
├── database/
│   ├── database.go       # Database operations
│   └── database_test.go  # Tests for database operations
├── handlers/
│   ├── tasks.go          # HTTP handlers for tasks
│   └── tasks_test.go     # Tests for HTTP handlers
├── models/
│   └── task.go           # Task data model
├── go.mod                # Go module file
└── README.md             # This file
```

## Features

- Complete CRUD operations for tasks
- RESTful API design
- SQLite database integration
- Comprehensive test suite with table-driven tests

## API Endpoints

- `GET /tasks` - Get all tasks
- `GET /tasks/{id}` - Get a specific task
- `POST /tasks` - Create a new task
- `PUT /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task

## How to Run

```bash
# Navigate to the project directory
cd task_manager_api

# Run the application
go run cmd/main.go
```

The server will start on port 8080. Use curl, Postman, or any HTTP client to interact with the API.

## Example API Requests

### Create a Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Complete Go assignment","description":"Finish the REST API project","status":"pending","due_date":"2025-06-10T00:00:00Z"}'
```

### Get All Tasks
```bash
curl http://localhost:8080/tasks
```

### Get a Specific Task
```bash
curl http://localhost:8080/tasks/1
```

### Update a Task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed"}'
```

### Delete a Task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests for a specific package
go test ./database
go test ./handlers
```

## Best Practices Demonstrated

1. **Code Organization**:
   - Separation of concerns with packages for models, database, and handlers
   - Clean API design with proper HTTP status codes

2. **Testing**:
   - Table-driven tests for comprehensive test coverage
   - In-memory database for testing
   - Tests for both database operations and HTTP handlers

3. **Error Handling**:
   - Proper error handling and reporting
   - Appropriate HTTP status codes for different scenarios

4. **Database**:
   - Safe database operations
   - Proper connection management
