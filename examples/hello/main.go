package main

import (
	"fmt"
	"net/http"

	"github.com/dayuoba/mulan"
)

func main() {
	fmt.Println("hello testing ")
	http.ListenAndServe(":9990", &mulan.OnConnect{})
}
