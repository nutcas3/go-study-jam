package database

import (
	"database/sql"
	"log"
	"task_manager_api/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the database connection
var DB *sql.DB

// InitDB initializes the database connection and creates the tasks table if it doesn't exist
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	createTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL,
		due_date DATETIME,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);`
	
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

// CreateTask adds a new task to the database
func CreateTask(task models.Task) (int64, error) {
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	
	if task.Status == "" {
		task.Status = "pending"
	}
	
	query := `INSERT INTO tasks 
		(title, description, status, due_date, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)`
	
	res, err := DB.Exec(query, 
		task.Title, 
		task.Description, 
		task.Status, 
		task.DueDate, 
		task.CreatedAt, 
		task.UpdatedAt)
	
	if err != nil {
		return 0, err
	}
	
	return res.LastInsertId()
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]models.Task, error) {
	query := `SELECT id, title, description, status, due_date, created_at, updated_at 
		FROM tasks ORDER BY created_at DESC`
	
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		var dueDate sql.NullTime
		
		err := rows.Scan(
			&task.ID, 
			&task.Title, 
			&task.Description, 
			&task.Status, 
			&dueDate, 
			&task.CreatedAt, 
			&task.UpdatedAt)
		
		if err != nil {
			return nil, err
		}
		
		if dueDate.Valid {
			task.DueDate = dueDate.Time
		}
		
		tasks = append(tasks, task)
	}
	
	return tasks, nil
}

// GetTaskByID retrieves a single task by ID
func GetTaskByID(id int) (models.Task, error) {
	query := `SELECT id, title, description, status, due_date, created_at, updated_at 
		FROM tasks WHERE id = ?`
	
	var task models.Task
	var dueDate sql.NullTime
	
	err := DB.QueryRow(query, id).Scan(
		&task.ID, 
		&task.Title, 
		&task.Description, 
		&task.Status, 
		&dueDate, 
		&task.CreatedAt, 
		&task.UpdatedAt)
	
	if dueDate.Valid {
		task.DueDate = dueDate.Time
	}
	
	return task, err
}

// UpdateTask updates an existing task
func UpdateTask(id int, task models.Task) error {
	existingTask, err := GetTaskByID(id)
	if err != nil {
		return err
	}
	
	// Update only the fields that are provided
	if task.Title != "" {
		existingTask.Title = task.Title
	}
	if task.Description != "" {
		existingTask.Description = task.Description
	}
	if task.Status != "" {
		existingTask.Status = task.Status
	}
	if !task.DueDate.IsZero() {
		existingTask.DueDate = task.DueDate
	}
	
	existingTask.UpdatedAt = time.Now()
	
	query := `UPDATE tasks SET 
		title = ?, 
		description = ?, 
		status = ?, 
		due_date = ?, 
		updated_at = ? 
		WHERE id = ?`
	
	_, err = DB.Exec(query, 
		existingTask.Title, 
		existingTask.Description, 
		existingTask.Status, 
		existingTask.DueDate, 
		existingTask.UpdatedAt, 
		id)
	
	return err
}

// DeleteTask removes a task from the database
func DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := DB.Exec(query, id)
	return err
}
