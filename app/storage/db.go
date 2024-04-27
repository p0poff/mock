package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
		CREATE TABLE IF NOT EXISTS setting (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			option VARCHAR(50) NOT NULL,
			value VARCHAR(500) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS route (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url VARCHAR(500) NOT NULL,
			method VARCHAR(12) NOT NULL,
			headers JSON NOT NULL DEFAULT "{}",
			body TEXT
		);
		INSERT INTO setting (option, value)
		VALUES ('default_headers', '{"Content-Type": "application/json"}')
	`

	// Execute the SQL query
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteDB) TableExists(tableName string) bool {
	var name string
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name=?;`
	err := s.db.QueryRow(query, tableName).Scan(&name)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("[CRITICAL] Query error: %s", err)
		}
		return false
	}
	return true
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
