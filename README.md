# Mulan is a super girl hero in China

## My Goal

* easy json
* easy router
* easy configuration
* no third deps
* high performance
* few footprint

## Roadmap

### V1.0.0

* cli tool for generate template
* json pre-define
* basic router

### V2.0.0

### V3.0.0

## Quick Start

```golang
package main

import ml "github.com/dayuoba/mulan"

func main() {

    server := ml.Server()

    router := server.Router()

    router.Use(func(c ml.Ctx, next ml.Next) {
        // do some stuff
        next()
    })

    router.Get("/", func(c ml.Ctx) {
        c.Send(ml.JSON{})
    })

    if err := server.Listen("8080"); err != {
        if err != nil {
            panic(err)
        }
        fmt.Println("server started at 8080")
    })
}

```
