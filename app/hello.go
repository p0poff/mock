package main

import "fmt"
import "github.com/jessevdk/go-flags"
import "github.com/p0poff/mock/app/storage"
import "os"
import "log"

var opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	FileDb  string `short:"f" long:"filedb" description:"file db path" required:"true"`

	Logger struct {
		FileName string `long:"file" env:"FILE_LOG"  default:"../data/hello.log" description:"location of spam log"`
	}
}

func main() {
	fmt.Println("Hello, World!")

	p := flags.NewParser(&opts, flags.Default)
	_, err := p.Parse()
	if err != nil {
		fmt.Println("Error parsing flags")
		return
	}

	//logger
	f, err := os.OpenFile(opts.Logger.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("[INFO] Hello, World!")

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

	err = db.CreateTables()
	if err != nil {
		fmt.Println("Error creating tables:", err)
		return
	}

	err = db.AddProduct("Apple", 0.5)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
}
