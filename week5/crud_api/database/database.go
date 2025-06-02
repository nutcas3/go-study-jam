package database

import (
	"crud_api/models"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// DB is the database connection
var DB *sql.DB

// InitDB initializes the database connection and creates the items table if it doesn't exist
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	createTable := `CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

// InsertItem adds a new item to the database
func InsertItem(name string) (int64, error) {
	res, err := DB.Exec("INSERT INTO items (name) VALUES (?)", name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetAllItems retrieves all items from the database
func GetAllItems() ([]models.Item, error) {
	rows, err := DB.Query("SELECT id, name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var items []models.Item
	for rows.Next() {
		var it models.Item
		if err := rows.Scan(&it.ID, &it.Name); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}

// GetItem retrieves a single item by ID
func GetItem(id int) (models.Item, error) {
	var item models.Item
	err := DB.QueryRow("SELECT id, name FROM items WHERE id = ?", id).Scan(&item.ID, &item.Name)
	return item, err
}

// UpdateItem updates an existing item
func UpdateItem(id int, name string) error {
	_, err := DB.Exec("UPDATE items SET name = ? WHERE id = ?", name, id)
	return err
}

// DeleteItem removes an item from the database
func DeleteItem(id int) error {
	_, err := DB.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}
