package main

import (
	"fmt"
	"github.com/p0poff/mock/app/circular_stack"
	"github.com/p0poff/mock/app/storage"
)

func main() {
	fmt.Println("Start")

	c := circular_stack.NewCircularStack(20)

	for i := 0; i < 200; i++ {
		c.Push(storage.Route{Id: i, Url: "pop"})
	}

	fmt.Println(c.All())

}
