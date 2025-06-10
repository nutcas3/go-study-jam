package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_manager_api/database"
	"task_manager_api/models"
	"testing"
	"time"
	"strconv"
)

// setupTest initializes the test environment
func setupTest(t *testing.T) {
	// Use an in-memory SQLite database for testing
	database.InitDB(":memory:")
}

// TestGetAllTasks tests the getAllTasks handler
func TestGetAllTasks(t *testing.T) {
	setupTest(t)

	// Create some test tasks
	tasks := []models.Task{
		{
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "pending",
		},
		{
			Title:       "Task 2",
			Description: "Description 2",
			Status:      "in_progress",
		},
	}

	for _, task := range tasks {
		_, err := database.CreateTask(task)
		if err != nil {
			t.Fatalf("Failed to create test task: %v", err)
		}
	}

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TasksHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var gotTasks []models.Task
	err = json.Unmarshal(rr.Body.Bytes(), &gotTasks)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify we got the expected number of tasks
	if len(gotTasks) != len(tasks) {
		t.Errorf("handler returned unexpected number of tasks: got %d want %d",
			len(gotTasks), len(tasks))
	}
}

// TestCreateTask tests the createTask handler
func TestCreateTask(t *testing.T) {
	setupTest(t)

	// Test cases
	testCases := []struct {
		name       string
		task       models.Task
		wantStatus int
	}{
		{
			name: "Valid Task",
			task: models.Task{
				Title:       "Test Task",
				Description: "This is a test task",
				Status:      "pending",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Empty Title",
			task: models.Task{
				Description: "Task with empty title",
				Status:      "pending",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert task to JSON
			taskJSON, err := json.Marshal(tc.task)
			if err != nil {
				t.Fatalf("Failed to marshal task: %v", err)
			}

			// Create a request with the task in the body
			req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(TasksHandler)

			// Call the handler
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.wantStatus)
			}

			// If success, check the response contains a task with an ID
			if tc.wantStatus == http.StatusCreated {
				var createdTask models.Task
				err = json.Unmarshal(rr.Body.Bytes(), &createdTask)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if createdTask.ID <= 0 {
					t.Errorf("Created task has invalid ID: %d", createdTask.ID)
				}

				if createdTask.Title != tc.task.Title {
					t.Errorf("Created task has wrong title: got %s want %s",
						createdTask.Title, tc.task.Title)
				}
			}
		})
	}
}

// TestGetTaskByID tests the getTaskByID handler
func TestGetTaskByID(t *testing.T) {
	setupTest(t)

	// Create a test task
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
	}

	id, err := database.CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}

	// Test cases
	strId := strconv.FormatInt(id, 10)
	testCases := []struct {
		name       string
		taskID     string
		wantStatus int
	}{
		{
			name:       "Existing Task",
			taskID:     "/tasks/" + strId,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Non-existent Task",
			taskID:     "/tasks/9999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Invalid ID",
			taskID:     "/tasks/invalid",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request
			req, err := http.NewRequest("GET", tc.taskID, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(TasksHandler)

			// Call the handler
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.wantStatus)
			}

			// If success, check the response contains the expected task
			if tc.wantStatus == http.StatusOK {
				var gotTask models.Task
				err = json.Unmarshal(rr.Body.Bytes(), &gotTask)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if gotTask.ID != int(id) {
					t.Errorf("Got task with wrong ID: got %d want %d",
						gotTask.ID, id)
				}

				if gotTask.Title != task.Title {
					t.Errorf("Got task with wrong title: got %s want %s",
						gotTask.Title, task.Title)
				}
			}
		})
	}
}

// TestUpdateTask tests the updateTask handler
func TestUpdateTask(t *testing.T) {
	setupTest(t)

	// Create a test task
	task := models.Task{
		Title:       "Original Title",
		Description: "Original Description",
		Status:      "pending",
	}

	id, err := database.CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}

	// Test cases
	strId := strconv.FormatInt(id, 10)
	testCases := []struct {
		name       string
		taskID     string
		updateTask models.Task
		wantStatus int
	}{
		{
			name:   "Update Title",
			taskID: "/tasks/" + strId,
			updateTask: models.Task{
				Title: "Updated Title",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "Update Status",
			taskID: "/tasks/" + strId,
			updateTask: models.Task{
				Status: "completed",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "Non-existent Task",
			taskID: "/tasks/9999",
			updateTask: models.Task{
				Title: "This Won't Work",
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "Invalid ID",
			taskID: "/tasks/invalid",
			updateTask: models.Task{
				Title: "This Won't Work Either",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert update task to JSON
			taskJSON, err := json.Marshal(tc.updateTask)
			if err != nil {
				t.Fatalf("Failed to marshal task: %v", err)
			}

			// Create a request with the task in the body
			req, err := http.NewRequest("PUT", tc.taskID, bytes.NewBuffer(taskJSON))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(TasksHandler)

			// Call the handler
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.wantStatus)
			}

			// If success, check the response contains the updated task
			if tc.wantStatus == http.StatusOK {
				var updatedTask models.Task
				err = json.Unmarshal(rr.Body.Bytes(), &updatedTask)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if tc.updateTask.Title != "" && updatedTask.Title != tc.updateTask.Title {
					t.Errorf("Updated task has wrong title: got %s want %s",
						updatedTask.Title, tc.updateTask.Title)
				}

				if tc.updateTask.Status != "" && updatedTask.Status != tc.updateTask.Status {
					t.Errorf("Updated task has wrong status: got %s want %s",
						updatedTask.Status, tc.updateTask.Status)
				}
			}
		})
	}
}

// TestDeleteTask tests the deleteTask handler
func TestDeleteTask(t *testing.T) {
	setupTest(t)

	// Create a test task
	task := models.Task{
		Title:  "Task to Delete",
		Status: "pending",
	}

	id, err := database.CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}

	// Test cases
	strId := strconv.FormatInt(id, 10)
	testCases := []struct {
		name       string
		taskID     string
		wantStatus int
	}{
		{
			name:       "Existing Task",
			taskID:     "/tasks/" + strId,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Non-existent Task",
			taskID:     "/tasks/9999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Invalid ID",
			taskID:     "/tasks/invalid",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request
			req, err := http.NewRequest("DELETE", tc.taskID, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(TasksHandler)

			// Call the handler
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.wantStatus)
			}

			// If success, verify the task was deleted
			if tc.wantStatus == http.StatusOK {
				// Try to get the task
				getReq, err := http.NewRequest("GET", tc.taskID, nil)
				if err != nil {
					t.Fatal(err)
				}

				getRr := httptest.NewRecorder()
				handler.ServeHTTP(getRr, getReq)

				// Should get a 404 Not Found
				if getRr.Code != http.StatusNotFound {
					t.Errorf("Task was not deleted, got status %d", getRr.Code)
				}
			}
		})
	}
}
