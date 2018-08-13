package main

import (
	"fmt"

	ml "github.com/dayuoba/mulan"
)

func main() {
	s := ml.Server()
	s.Use(func(c *ml.Context, next ml.Next) {
		fmt.Println("setting up middleware")
	})
	s.Listen(":9999")
}
