package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func NewSqliteDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveTmpDB(file multipart.File, handler *multipart.FileHeader) (string, error) {
	path := fmt.Sprintf("/tmp/%s", handler.Filename)
	tempFile, err := os.Create(path)
	if err != nil {
		log.Printf("[ERROR] Error creating the tmp file")
		return path, fmt.Errorf("Error creating the tmp file")
	}
	defer tempFile.Close()

	// Write the content from the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		log.Printf("[ERROR] Error saving the tmp file")
		return path, fmt.Errorf("Error saving the tmp file")
	}

	return path, nil
}
