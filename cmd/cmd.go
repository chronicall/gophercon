package main

import (
    "log"
    "os"

    "github.com/chronicall/gophercon/pkg/routing"
    "github.com/chronicall/gophercon/pkg/webserver"
	"github.com/chronicall/gophercon/version"
)

// go run ./cmd/cmd.go
// curl -i http://127.0.0.1:8000/home
func main() {
    log.Printf(
		"Service is starting.\nVersion is %s, commit is %s, time is %s.",
		version.Release, version.Commit, version.BuildTime,
	)

    // You can also use github.com/kelseyhightower/envconfig
    // to keep your config more structured
    port := os.Getenv("PORT")
    if len(port) == 0 {
        log.Fatal("Service port wasn't set.")
    }

    r := routing.BaseRouter()
    ws := webserver.New("", port, r)

    log.Fatal(ws.Start())
}
