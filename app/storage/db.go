package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(db *sql.DB) (*SQLiteDB, error) {
	return &SQLiteDB{
		db: db,
	}, nil
}

func (s *SQLiteDB) Close() error {
	return s.db.Close()
}

func (s *SQLiteDB) CreateTables() error {
	// SQL query to create tables
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		);
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price REAL NOT NULL
		);
	`

	// Execute the SQL query
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteDB) AddProduct(name string, price float64) error {
	query := `
		INSERT INTO products (name, price)
		VALUES (?, ?)
	`
	_, err := s.db.Exec(query, name, price)
	if err != nil {
		return err
	}
	return nil
}

// Add other methods for interacting with the SQLite database here
