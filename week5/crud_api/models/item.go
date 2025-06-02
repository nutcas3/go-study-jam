package models

// Item represents a basic item in our CRUD application
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
