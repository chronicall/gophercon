package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

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

    shutdown := make(chan error, 2)

    // You can also use github.com/kelseyhightower/envconfig
    // to keep your config more structured

    // Main web server
    port := os.Getenv("PORT")
    if len(port) == 0 {
        log.Fatal("Service port wasn't set.")
    }
    r := routing.BaseRouter()
    ws := webserver.New("", port, r)
    go func() {
        err := ws.Start()
        shutdown <-err
    }()

    // Diagnostics server
    internalPort := os.Getenv("INTERNAL_PORT")
    if len(internalPort) == 0 {
        log.Fatal("Internal port wasn't set.")
    }
    diagnosticsRouter := routing.DiagnosticsRouter()
    diagnosticsServer := webserver.New("", internalPort, diagnosticsRouter)
    go func() {
        err := diagnosticsServer.Start()
        shutdown <-err
    }()

    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

    select {
    case killSignal := <-interrupt:
        log.Printf("Got %s. Stopping...", killSignal)
    case err := <-shutdown:
        log.Printf("Got an error: '%s'. Stopping...", err)
    }

    err := ws.Stop()
    if err != nil {
        log.Print(err)
    }

    err = diagnosticsServer.Stop()
    if err != nil {
        log.Print(err)
    }

    // stop extra tasks

    log.Print("Service was stopped.")
}
