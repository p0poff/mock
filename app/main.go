package main

import "fmt"
import "os"
import "log"
import "github.com/jessevdk/go-flags"
import "github.com/p0poff/mock/app/storage"
import "github.com/p0poff/mock/app/server"

type opts struct {
	FileDb  string `short:"f" long:"filedb" env:"FILE_DB" description:"file db path db" required:"true"`
	FileLog string `short:"l" long:"filelog" env:"FILE_LOG"  description:"location of spam log" required:"true"`
	Port    string `short:"p" long:"port" env:"PORT" default:"8080" description:"port to listen on"`
}

func main() {
	var opts opts

	//options
	p := flags.NewParser(&opts, flags.Default)
	_, err := p.Parse()

	if err != nil {
		fmt.Println("Error parsing flags")
		return
	}

	//logger
	f, err := os.OpenFile(opts.FileLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println("[INFO] App start!")

	//DB
	sqlite, err := storage.NewSqliteDB(opts.FileDb)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	db, err := storage.NewSQLiteDB(sqlite)
	if err != nil {
		fmt.Println("Error creating SQLiteDB:", err)
		return
	}

	defer db.Close()

	if !db.TableExists("route") {
		fmt.Println("Creating tables")
		err = db.CreateTables()
		if err != nil {
			fmt.Println("Error creating tables:", err)
			return
		}
	}

	//server
	s := server.NewServer(opts.Port, db)

	if err = s.Start(); err != nil {
		log.Printf("[ERROR] %v", err)
	}

	log.Println("[INFO] App stop!")
}
