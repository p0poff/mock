package storage

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SQLiteDB struct {
	db *sql.DB
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func NewSQLiteDB(db *sql.DB) (*SQLiteDB, error) {
	return &SQLiteDB{
		db: db,
	}, nil
}

func getRouteFromDb(row scanner) (Route, error) {
	var route Route
	var headersJSON string
	err := row.Scan(&route.Id, &route.Url, &route.Method, &headersJSON, &route.StatusCode, &route.Body)
	if err != nil {
		log.Println("[ERROR] Failed to scan route (GetRoute):", err)
		return route, err
	}

	err = json.Unmarshal([]byte(headersJSON), &route.Headers)
	if err != nil {
		log.Println("[ERROR] Failed to unmarshal headers (GetRoute):", err)
		return route, err
	}

	return route, nil
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
			status_code INTEGER NOT NULL DEFAULT 200,
			body TEXT
		);
		INSERT INTO setting (option, value)
		VALUES ('default_headers', '{"Content-Type": "application/json"}');
	
		CREATE UNIQUE INDEX IF NOT EXISTS idx_route_url_method ON route (url, method);
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

func (s *SQLiteDB) AddRoute(route Route) error {
	query := `
		INSERT INTO route (url, method, headers, status_code, body)
		VALUES (?, ?, ?, ?, ?)
	`
	headersJSON, err := json.Marshal(route.Headers)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(query, route.Url, route.Method, headersJSON, route.StatusCode, route.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteDB) EditRoute(route Route) error {
	query := `
		UPDATE route
		SET url = ?, method = ?, headers = ?, status_code = ?, body = ?
		WHERE id = ?
	`
	headersJSON, err := json.Marshal(route.Headers)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(query, route.Url, route.Method, headersJSON, route.StatusCode, route.Body, route.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteDB) DeleteRoute(route Route) error {
	query := `
		DELETE FROM route
		WHERE id = ?
	`
	_, err := s.db.Exec(query, route.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteDB) GetRoutes() ([]Route, error) {
	query := `
		SELECT id, url, method, headers, status_code, body
		FROM route
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []Route
	for rows.Next() {
		route, err := getRouteFromDb(rows)
		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	return routes, nil
}

func (s *SQLiteDB) GetRoute(url string, method string) (Route, error) {
	query := `
		SELECT id, url, method, headers, status_code, body
		FROM route
		WHERE url = ? AND method = ?
	`
	row := s.db.QueryRow(query, url, method)

	route, err := getRouteFromDb(row)
	if err != nil {
		return route, err
	}

	return route, nil

}

// Add other methods for interacting with the SQLite database here
