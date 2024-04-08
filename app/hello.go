package main

import "fmt"
import "github.com/jessevdk/go-flags"

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

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
}
