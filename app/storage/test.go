package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func TestSqlConnect() *sql.DB {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "/home/pop/pet/go/mock/data/test.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil
	}
	defer db.Close()

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	fmt.Println("Connected to SQLite database!")

	return db
}
