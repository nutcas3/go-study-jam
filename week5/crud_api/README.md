# REST API for CRUD Application

This project is part of Week 5 of the Go course. It is a REST API for a CRUD application with unit tests.

## How to run

Run the CRUD API server from the `cmd` directory:

```bash
# Example:
go run cmd/main.go
```

The server will start on port 8080. Use curl or Postman to interact with the API endpoints (e.g., GET/POST http://localhost:8080/items).

## Tasks
- Implement basic CRUD operations (Create, Read, Update, Delete).
- Use the database/sql package to connect to a SQL database.
- Write unit and table-driven tests.
