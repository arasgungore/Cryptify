package database

import "fmt"

// Database simulates a simple in-memory database
type Database struct {
	data map[string]interface{}
}

// NewDatabase creates a new instance of Database
func NewDatabase() *Database {
	return &Database{
		data: make(map[string]interface{}),
	}
}

// Save saves data to the database
func (db *Database) Save(key string, value interface{}) {
	db.data[key] = value
	fmt.Printf("Data saved: %v\n", value)
}

// Get retrieves data from the database
func (db *Database) Get(key string) (interface{}, bool) {
	value, exists := db.data[key]
	return value, exists
}
