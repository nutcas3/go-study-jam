# Task Manager API Testing Guide

This document provides an overview of the testing strategies and techniques implemented in the Task Manager API project. The tests demonstrate various Go testing approaches including unit tests, integration tests, table-driven tests, benchmark tests, and edge case testing.

## Testing Structure

The testing suite is organized into several files, each focusing on different aspects of testing:

1. **Unit Tests**:
   - `database/database_test.go`: Tests for database operations
   - `handlers/tasks_test.go`: Tests for HTTP handlers

2. **Integration Tests**:
   - `tests/api_test.go`: End-to-end API tests covering full task lifecycle

3. **Advanced Testing Techniques**:
   - `tests/table_driven_test.go`: Table-driven tests for API endpoints
   - `tests/benchmark_test.go`: Performance benchmarks for API operations
   - `tests/edge_cases_test.go`: Tests for edge cases and error handling

## Testing Approaches

### 1. Unit Testing

Unit tests focus on testing individual components in isolation:

- **Database Tests**: Test CRUD operations directly against the database layer
- **Handler Tests**: Test HTTP handlers with mocked requests and responses

Key techniques:
- Using in-memory SQLite database for testing
- Isolating components for focused testing
- Testing both success and error paths

### 2. Table-Driven Tests

Table-driven tests allow testing multiple scenarios with different inputs and expected outputs using a single test function. This approach:

- Reduces code duplication
- Makes it easy to add new test cases
- Provides clear documentation of test scenarios

Example from `table_driven_test.go`:
```go
testCases := []struct {
    name           string
    task           models.Task
    expectedStatus int
    validateFunc   func(*testing.T, models.Task)
}{
    {
        name: "Valid Task",
        task: models.Task{
            Title:       "Valid Task",
            Description: "This is a valid task",
            Status:      "pending",
            DueDate:     time.Now().Add(24 * time.Hour),
        },
        expectedStatus: http.StatusCreated,
        validateFunc: func(t *testing.T, task models.Task) {
            if task.ID <= 0 {
                t.Errorf("Expected valid ID, got %d", task.ID)
            }
            // More validations...
        },
    },
    // More test cases...
}
```

### 3. Benchmark Testing

Benchmark tests measure the performance of operations:

- **Create Task**: Measures task creation performance
- **Get All Tasks**: Measures retrieving all tasks
- **Get Task By ID**: Measures retrieving a single task
- **Update Task**: Measures task update performance
- **Delete Task**: Measures task deletion performance
- **CRUD Operations**: Measures a complete CRUD cycle

Results are reported in nanoseconds per operation, providing insights into API performance.

### 4. Edge Case Testing

Edge case tests verify the API's behavior in unusual or extreme situations:

- **Invalid Requests**: Testing with malformed JSON, invalid IDs, etc.
- **Concurrent Requests**: Testing behavior under concurrent load
- **Data Validation**: Testing with extremely large payloads, special characters, etc.
- **Error Recovery**: Testing the API's ability to recover from errors

### 5. Integration Testing

Integration tests verify that all components work together correctly:

- **Task Lifecycle**: Tests the complete lifecycle of a task (create, read, update, delete)
- **Error Handling**: Tests how the API handles errors in a real-world scenario

## Running Tests

To run all tests:

```bash
go test ./... -v
```

To run specific test files:

```bash
go test ./tests/table_driven_test.go -v
```

To run benchmark tests:

```bash
go test ./tests/benchmark_test.go -bench=.
```

## Best Practices Demonstrated

1. **Test Independence**: Each test can run independently without relying on other tests
2. **Clean Setup/Teardown**: Using in-memory databases and proper test cleanup
3. **Comprehensive Coverage**: Testing both happy paths and error cases
4. **Readable Tests**: Clear test names and organization
5. **Performance Testing**: Benchmark tests to identify bottlenecks
6. **Edge Case Testing**: Testing unusual inputs and scenarios
7. **Table-Driven Tests**: Efficient testing of multiple scenarios
8. **Separation of Concerns**: Unit tests for individual components, integration tests for the whole system

## Conclusion

The testing suite demonstrates a comprehensive approach to testing a Go REST API, covering all aspects from unit testing to performance benchmarking. These testing strategies ensure the reliability, performance, and correctness of the Task Manager API.
