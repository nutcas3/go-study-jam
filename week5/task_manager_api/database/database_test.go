package database

import (
	"os"
	"task_manager_api/models"
	"testing"
	"time"
)

// Setup test database
func setupTestDB(t *testing.T) {
	// Use an in-memory SQLite database for testing
	InitDB(":memory:")
}

// TestCreateTask tests the CreateTask function
func TestCreateTask(t *testing.T) {
	setupTestDB(t)

	// Test cases
	testCases := []struct {
		name     string
		task     models.Task
		wantErr  bool
		checkID  bool
	}{
		{
			name: "Valid Task",
			task: models.Task{
				Title:       "Test Task",
				Description: "This is a test task",
				Status:      "pending",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			wantErr: false,
			checkID: true,
		},
		{
			name: "Empty Title",
			task: models.Task{
				Description: "Task with empty title",
				Status:      "pending",
			},
			wantErr: false, // SQLite doesn't enforce NOT NULL at driver level
			checkID: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := CreateTask(tc.task)
			
			// Check error
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// Check ID
			if tc.checkID && id <= 0 {
				t.Errorf("CreateTask() returned invalid ID: %d", id)
			}
		})
	}
}

// TestGetAllTasks tests the GetAllTasks function
func TestGetAllTasks(t *testing.T) {
	setupTestDB(t)
	
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
		_, err := CreateTask(task)
		if err != nil {
			t.Fatalf("Failed to create test task: %v", err)
		}
	}
	
	// Test GetAllTasks
	gotTasks, err := GetAllTasks()
	if err != nil {
		t.Errorf("GetAllTasks() error = %v", err)
		return
	}
	
	// Check if we got the expected number of tasks
	if len(gotTasks) != len(tasks) {
		t.Errorf("GetAllTasks() returned %d tasks, want %d", len(gotTasks), len(tasks))
	}
}

// TestGetTaskByID tests the GetTaskByID function
func TestGetTaskByID(t *testing.T) {
	setupTestDB(t)
	
	// Create a test task
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
	}
	
	id, err := CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}
	
	// Test cases
	testCases := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "Existing Task",
			id:      int(id),
			wantErr: false,
		},
		{
			name:    "Non-existent Task",
			id:      9999,
			wantErr: true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotTask, err := GetTaskByID(tc.id)
			
			// Check error
			if (err != nil) != tc.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// If we expect success, check the task details
			if !tc.wantErr {
				if gotTask.ID != tc.id {
					t.Errorf("GetTaskByID() got task with ID = %d, want %d", gotTask.ID, tc.id)
				}
				if gotTask.Title != task.Title {
					t.Errorf("GetTaskByID() got task with Title = %s, want %s", gotTask.Title, task.Title)
				}
			}
		})
	}
}

// TestUpdateTask tests the UpdateTask function
func TestUpdateTask(t *testing.T) {
	setupTestDB(t)
	
	// Create a test task
	task := models.Task{
		Title:       "Original Title",
		Description: "Original Description",
		Status:      "pending",
	}
	
	id, err := CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}
	
	// Test cases
	testCases := []struct {
		name       string
		id         int
		updateTask models.Task
		wantErr    bool
	}{
		{
			name: "Update Title",
			id:   int(id),
			updateTask: models.Task{
				Title: "Updated Title",
			},
			wantErr: false,
		},
		{
			name: "Update Status",
			id:   int(id),
			updateTask: models.Task{
				Status: "completed",
			},
			wantErr: false,
		},
		{
			name: "Non-existent Task",
			id:   9999,
			updateTask: models.Task{
				Title: "This Won't Work",
			},
			wantErr: true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := UpdateTask(tc.id, tc.updateTask)
			
			// Check error
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// If we expect success, check the task was updated
			if !tc.wantErr {
				updatedTask, err := GetTaskByID(tc.id)
				if err != nil {
					t.Errorf("Failed to get updated task: %v", err)
					return
				}
				
				// Check if the field was updated
				if tc.updateTask.Title != "" && updatedTask.Title != tc.updateTask.Title {
					t.Errorf("UpdateTask() failed to update Title, got %s, want %s", 
						updatedTask.Title, tc.updateTask.Title)
				}
				if tc.updateTask.Status != "" && updatedTask.Status != tc.updateTask.Status {
					t.Errorf("UpdateTask() failed to update Status, got %s, want %s", 
						updatedTask.Status, tc.updateTask.Status)
				}
			}
		})
	}
}

// TestDeleteTask tests the DeleteTask function
func TestDeleteTask(t *testing.T) {
	setupTestDB(t)
	
	// Create a test task
	task := models.Task{
		Title:  "Task to Delete",
		Status: "pending",
	}
	
	id, err := CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create test task: %v", err)
	}
	
	// Test cases
	testCases := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "Existing Task",
			id:      int(id),
			wantErr: false,
		},
		{
			name:    "Non-existent Task",
			id:      9999,
			wantErr: false, // SQLite doesn't return error for non-existent ID
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := DeleteTask(tc.id)
			
			// Check error
			if (err != nil) != tc.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// Verify task was deleted
			if !tc.wantErr {
				_, err := GetTaskByID(tc.id)
				if err == nil {
					t.Errorf("DeleteTask() failed to delete task with ID = %d", tc.id)
				}
			}
		})
	}
}

// TestMain handles setup and teardown for all tests
func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	
	// Clean up
	os.Exit(code)
}
