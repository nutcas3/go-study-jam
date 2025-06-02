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

// setupTableTestServer creates a test server for table-driven tests
func setupTableTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/tasks" || r.URL.Path == "/tasks/" || r.URL.Path[:7] == "/tasks/" {
			handlers.TasksHandler(w, r)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
}

// TestCreateTaskTable demonstrates table-driven testing for task creation
func TestCreateTaskTable(t *testing.T) {
	// Initialize the database
	database.InitDB(":memory:")
	
	// Set up the test server
	server := setupTableTestServer()
	defer server.Close()
	
	// Define test cases
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
				if task.Title != "Valid Task" {
					t.Errorf("Expected title 'Valid Task', got '%s'", task.Title)
				}
				if task.Status != "pending" {
					t.Errorf("Expected status 'pending', got '%s'", task.Status)
				}
			},
		},
		{
			name: "Empty Title",
			task: models.Task{
				Description: "Task with empty title",
				Status:      "pending",
			},
			expectedStatus: http.StatusBadRequest,
			validateFunc:   nil, // No validation needed for error cases
		},
		{
			name: "Very Long Title",
			task: models.Task{
				Title:       string(make([]byte, 1000)), // 1000 character title
				Description: "Task with very long title",
				Status:      "pending",
			},
			expectedStatus: http.StatusCreated, // Should still work
			validateFunc: func(t *testing.T, task models.Task) {
				if task.ID <= 0 {
					t.Errorf("Expected valid ID, got %d", task.ID)
				}
				if len(task.Title) != 1000 {
					t.Errorf("Expected title length 1000, got %d", len(task.Title))
				}
			},
		},
		{
			name: "Invalid Status",
			task: models.Task{
				Title:       "Task with invalid status",
				Description: "This task has an invalid status",
				Status:      "invalid_status",
			},
			expectedStatus: http.StatusCreated, // Should still work, we don't validate status
			validateFunc: func(t *testing.T, task models.Task) {
				if task.Status != "invalid_status" {
					t.Errorf("Expected status 'invalid_status', got '%s'", task.Status)
				}
			},
		},
		{
			name: "Past Due Date",
			task: models.Task{
				Title:       "Task with past due date",
				Description: "This task has a due date in the past",
				Status:      "pending",
				DueDate:     time.Now().Add(-24 * time.Hour),
			},
			expectedStatus: http.StatusCreated, // Should still work, we don't validate due date
			validateFunc: func(t *testing.T, task models.Task) {
				if task.DueDate.After(time.Now()) {
					t.Errorf("Expected due date in the past, got future date")
				}
			},
		},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert task to JSON
			taskJSON, err := json.Marshal(tc.task)
			if err != nil {
				t.Fatalf("Failed to marshal task: %v", err)
			}
			
			// Send POST request to create task
			resp, err := http.Post(
				server.URL+"/tasks",
				"application/json",
				bytes.NewBuffer(taskJSON),
			)
			if err != nil {
				t.Fatalf("Failed to create task: %v", err)
			}
			defer resp.Body.Close()
			
			// Check status code
			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, resp.StatusCode)
			}
			
			// If we expect success and have a validation function, validate the response
			if tc.expectedStatus == http.StatusCreated && tc.validateFunc != nil {
				var createdTask models.Task
				if err := json.NewDecoder(resp.Body).Decode(&createdTask); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				
				// Run the validation function
				tc.validateFunc(t, createdTask)
			}
		})
	}
}

// TestUpdateTaskTable demonstrates table-driven testing for task updates
func TestUpdateTaskTable(t *testing.T) {
	// Initialize the database
	database.InitDB(":memory:")
	
	// Set up the test server
	server := setupTableTestServer()
	defer server.Close()
	
	// Create a test task to update
	task := models.Task{
		Title:       "Original Task",
		Description: "Original description",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
	}
	
	// Create the task
	taskJSON, _ := json.Marshal(task)
	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewBuffer(taskJSON),
	)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}
	
	// Get the created task ID
	var createdTask models.Task
	json.NewDecoder(resp.Body).Decode(&createdTask)
	resp.Body.Close()
	
	// Define test cases for updates
	testCases := []struct {
		name           string
		update         models.Task
		expectedStatus int
		validateFunc   func(*testing.T, models.Task)
	}{
		{
			name: "Update Title",
			update: models.Task{
				Title: "Updated Title",
			},
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, task models.Task) {
				if task.Title != "Updated Title" {
					t.Errorf("Expected title 'Updated Title', got '%s'", task.Title)
				}
				// Description should remain unchanged
				if task.Description != "Original description" {
					t.Errorf("Expected description to remain unchanged, got '%s'", task.Description)
				}
			},
		},
		{
			name: "Update Status",
			update: models.Task{
				Status: "completed",
			},
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, task models.Task) {
				if task.Status != "completed" {
					t.Errorf("Expected status 'completed', got '%s'", task.Status)
				}
				// Title should remain from previous update
				if task.Title != "Updated Title" {
					t.Errorf("Expected title to remain 'Updated Title', got '%s'", task.Title)
				}
			},
		},
		{
			name: "Update Description",
			update: models.Task{
				Description: "New description",
			},
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, task models.Task) {
				if task.Description != "New description" {
					t.Errorf("Expected description 'New description', got '%s'", task.Description)
				}
				// Status should remain from previous update
				if task.Status != "completed" {
					t.Errorf("Expected status to remain 'completed', got '%s'", task.Status)
				}
			},
		},
		{
			name: "Update Due Date",
			update: models.Task{
				DueDate: time.Now().Add(48 * time.Hour),
			},
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, task models.Task) {
				// Due date should be updated, but it's hard to compare exact times
				// So we'll just check it's in the future
				if task.DueDate.Before(time.Now()) {
					t.Errorf("Expected due date in the future, got past date")
				}
			},
		},
		{
			name: "Update Multiple Fields",
			update: models.Task{
				Title:       "Final Title",
				Description: "Final description",
				Status:      "in_progress",
			},
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, task models.Task) {
				if task.Title != "Final Title" {
					t.Errorf("Expected title 'Final Title', got '%s'", task.Title)
				}
				if task.Description != "Final description" {
					t.Errorf("Expected description 'Final description', got '%s'", task.Description)
				}
				if task.Status != "in_progress" {
					t.Errorf("Expected status 'in_progress', got '%s'", task.Status)
				}
			},
		},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert update to JSON
			updateJSON, err := json.Marshal(tc.update)
			if err != nil {
				t.Fatalf("Failed to marshal update: %v", err)
			}
			
			// Create a PUT request
			req, err := http.NewRequest(
				http.MethodPut,
				fmt.Sprintf("%s/tasks/%d", server.URL, createdTask.ID),
				bytes.NewBuffer(updateJSON),
			)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			
			// Send the request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to update task: %v", err)
			}
			defer resp.Body.Close()
			
			// Check status code
			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, resp.StatusCode)
			}
			
			// If we expect success and have a validation function, validate the response
			if tc.expectedStatus == http.StatusOK && tc.validateFunc != nil {
				var updatedTask models.Task
				if err := json.NewDecoder(resp.Body).Decode(&updatedTask); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				
				// Run the validation function
				tc.validateFunc(t, updatedTask)
			}
		})
	}
}
