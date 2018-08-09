package main

import "github.com/dayuoba/mulan"
import "net/http"

func main() {
	oc := &mulan.OnConnect{}
	http.ListenAndServe(":9992", oc)
}
