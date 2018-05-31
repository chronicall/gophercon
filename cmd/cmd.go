package main

import (
    "log"
    "net/http"

    "github.com/chronicall/gophercon/pkg/routing"
)

// go run ./cmd/cmd.go
// curl -i http://127.0.0.1:8000/home
func main() {
    log.Printf("Service is starting....")

    r := routing.BaseRouter()

    http.ListenAndServe(":8000", r)
}
