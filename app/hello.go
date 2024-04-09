package main

import "fmt"
import "github.com/jessevdk/go-flags"
import "github.com/p0poff/mock/app/storage"

var opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

func main() {
	fmt.Println("Hello, World!")

	p := flags.NewParser(&opts, flags.Default)
	_, err := p.Parse()
	if err != nil {
		fmt.Println("Error parsing flags")
		return
	}

	db := storage.TestSqlConnect()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		if err.Error() == "sql: database is closed" {
			fmt.Println("must be restore db connection")
		}
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
}
