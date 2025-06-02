package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_manager_api/database"
	"task_manager_api/handlers"
	"task_manager_api/models"
	"testing"
	"time"
)

// setupBenchmarkServer creates a test server for benchmarks
func setupBenchmarkServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/tasks" || r.URL.Path == "/tasks/" || r.URL.Path[:7] == "/tasks/" {
			handlers.TasksHandler(w, r)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
}

// setupBenchmarkDB initializes the database for benchmarks
func setupBenchmarkDB() {
	database.InitDB(":memory:")
}

// BenchmarkCreateTask benchmarks the task creation process
func BenchmarkCreateTask(b *testing.B) {
	// Initialize the database
	setupBenchmarkDB()
	
	// Set up the test server
	server := setupBenchmarkServer()
	defer server.Close()
	
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create a task with a unique title for each iteration
		task := models.Task{
			Title:       fmt.Sprintf("Benchmark Task %d", i),
			Description: "This is a benchmark task",
			Status:      "pending",
			DueDate:     time.Now().Add(24 * time.Hour),
		}
		
		// Convert task to JSON
		taskJSON, err := json.Marshal(task)
		if err != nil {
			b.Fatalf("Failed to marshal task: %v", err)
		}
		
		// Send POST request to create task
		resp, err := http.Post(
			server.URL+"/tasks",
			"application/json",
			bytes.NewBuffer(taskJSON),
		)
		if err != nil {
			b.Fatalf("Failed to create task: %v", err)
		}
		
		// Check status code
		if resp.StatusCode != http.StatusCreated {
			b.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
		}
		
		resp.Body.Close()
	}
}

// BenchmarkGetAllTasks benchmarks retrieving all tasks
func BenchmarkGetAllTasks(b *testing.B) {
	// Initialize the database
	setupBenchmarkDB()
	
	// Set up the test server
	server := setupBenchmarkServer()
	defer server.Close()
	
	// Create some test tasks
	numTasks := 100
	for i := 0; i < numTasks; i++ {
		task := models.Task{
			Title:       fmt.Sprintf("Benchmark Task %d", i),
			Description: "This is a benchmark task",
			Status:      "pending",
		}
		
		taskJSON, _ := json.Marshal(task)
		resp, err := http.Post(
			server.URL+"/tasks",
			"application/json",
			bytes.NewBuffer(taskJSON),
		)
		if err != nil {
			b.Fatalf("Failed to create test task: %v", err)
		}
		resp.Body.Close()
	}
	
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Send GET request to retrieve all tasks
		resp, err := http.Get(server.URL + "/tasks")
		if err != nil {
			b.Fatalf("Failed to get tasks: %v", err)
		}
		
		// Check status code
		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		
		// Parse response
		var tasks []models.Task
		if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
			b.Fatalf("Failed to decode response: %v", err)
		}
		
		// Verify we got the expected number of tasks
		if len(tasks) != numTasks {
			b.Fatalf("Expected %d tasks, got %d", numTasks, len(tasks))
		}
		
		resp.Body.Close()
	}
}

// BenchmarkGetTaskByID benchmarks retrieving a single task
func BenchmarkGetTaskByID(b *testing.B) {
	// Initialize the database
	setupBenchmarkDB()
	
	// Set up the test server
	server := setupBenchmarkServer()
	defer server.Close()
	
	// Create a test task
	task := models.Task{
		Title:       "Benchmark Task",
		Description: "This is a benchmark task",
		Status:      "pending",
	}
	
	taskJSON, _ := json.Marshal(task)
	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewBuffer(taskJSON),
	)
	if err != nil {
		b.Fatalf("Failed to create test task: %v", err)
	}
	
	// Get the created task ID
	var createdTask models.Task
	json.NewDecoder(resp.Body).Decode(&createdTask)
	resp.Body.Close()
	
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Send GET request to retrieve the task
		resp, err := http.Get(fmt.Sprintf("%s/tasks/%d", server.URL, createdTask.ID))
		if err != nil {
			b.Fatalf("Failed to get task: %v", err)
		}
		
		// Check status code
		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		
		resp.Body.Close()
	}
}

// BenchmarkUpdateTask benchmarks updating a task
func BenchmarkUpdateTask(b *testing.B) {
	// Initialize the database
	setupBenchmarkDB()
	
	// Set up the test server
	server := setupBenchmarkServer()
	defer server.Close()
	
	// Create a test task
	task := models.Task{
		Title:       "Benchmark Task",
		Description: "This is a benchmark task",
		Status:      "pending",
	}
	
	taskJSON, _ := json.Marshal(task)
	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewBuffer(taskJSON),
	)
	if err != nil {
		b.Fatalf("Failed to create test task: %v", err)
	}
	
	// Get the created task ID
	var createdTask models.Task
	json.NewDecoder(resp.Body).Decode(&createdTask)
	resp.Body.Close()
	
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create an update with a unique title for each iteration
		update := models.Task{
			Title: fmt.Sprintf("Updated Title %d", i),
		}
		
		// Convert update to JSON
		updateJSON, err := json.Marshal(update)
		if err != nil {
			b.Fatalf("Failed to marshal update: %v", err)
		}
		
		// Create a PUT request
		req, err := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/tasks/%d", server.URL, createdTask.ID),
			bytes.NewBuffer(updateJSON),
		)
		if err != nil {
			b.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		
		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			b.Fatalf("Failed to update task: %v", err)
		}
		
		// Check status code
		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		
		resp.Body.Close()
	}
}

// BenchmarkDeleteTask benchmarks deleting a task
func BenchmarkDeleteTask(b *testing.B) {
	// Initialize the database
	setupBenchmarkDB()
	
	// Set up the test server
	server := setupBenchmarkServer()
	defer server.Close()
	
	// Create tasks to delete
	taskIDs := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		task := models.Task{
			Title:       fmt.Sprintf("Task to Delete %d", i),
			Description: "This task will be deleted",
			Status:      "pending",
		}
		
		taskJSON, _ := json.Marshal(task)
		resp, err := http.Post(
			server.URL+"/tasks",
			"application/json",
			bytes.NewBuffer(taskJSON),
		)
		if err != nil {
			b.Fatalf("Failed to create test task: %v", err)
		}
		
		// Get the created task ID
		var createdTask models.Task
		json.NewDecoder(resp.Body).Decode(&createdTask)
		resp.Body.Close()
		
		taskIDs[i] = createdTask.ID
	}
	
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create a DELETE request
		req, err := http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/tasks/%d", server.URL, taskIDs[i]),
			nil,
		)
		if err != nil {
			b.Fatalf("Failed to create request: %v", err)
		}
		
		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			b.Fatalf("Failed to delete task: %v", err)
		}
		
		// Check status code
		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		
		resp.Body.Close()
	}
}

// BenchmarkCRUDOperations benchmarks a complete CRUD cycle
func BenchmarkCRUDOperations(b *testing.B) {
	// Initialize the database
	setupBenchmarkDB()
	
	// Set up the test server
	server := setupBenchmarkServer()
	defer server.Close()
	
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Step 1: Create a task
		task := models.Task{
			Title:       fmt.Sprintf("CRUD Task %d", i),
			Description: "This is a CRUD benchmark task",
			Status:      "pending",
		}
		
		taskJSON, _ := json.Marshal(task)
		resp, err := http.Post(
			server.URL+"/tasks",
			"application/json",
			bytes.NewBuffer(taskJSON),
		)
		if err != nil {
			b.Fatalf("Failed to create task: %v", err)
		}
		
		// Get the created task ID
		var createdTask models.Task
		json.NewDecoder(resp.Body).Decode(&createdTask)
		resp.Body.Close()
		
		// Step 2: Get the task
		resp, err = http.Get(fmt.Sprintf("%s/tasks/%d", server.URL, createdTask.ID))
		if err != nil {
			b.Fatalf("Failed to get task: %v", err)
		}
		resp.Body.Close()
		
		// Step 3: Update the task
		update := models.Task{
			Title:  fmt.Sprintf("Updated CRUD Task %d", i),
			Status: "in_progress",
		}
		
		updateJSON, _ := json.Marshal(update)
		req, _ := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/tasks/%d", server.URL, createdTask.ID),
			bytes.NewBuffer(updateJSON),
		)
		req.Header.Set("Content-Type", "application/json")
		
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			b.Fatalf("Failed to update task: %v", err)
		}
		resp.Body.Close()
		
		// Step 4: Delete the task
		req, _ = http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/tasks/%d", server.URL, createdTask.ID),
			nil,
		)
		
		resp, err = client.Do(req)
		if err != nil {
			b.Fatalf("Failed to delete task: %v", err)
		}
		resp.Body.Close()
	}
}
