package main

import (
    "fmt"
    "net/http"
)

func main() {
    err := http.ListenAndServe(":9991", nil)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}
