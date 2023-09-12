package main

import (
    "fmt"
    "net/http"
    "time"
    "os"

    "github.com/LaMoldy/groupme/server/pkg/router"
    _ "github.com/lib/pq"
)

func main() {
    // Gets the configured setup setup router
    router := router.SetupRouter()

    // Configures the server
    server := &http.Server {
        Addr:         ":8080",
        Handler:      router,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // Sets up the servers listener
    err := server.ListenAndServe()
    if err != nil {
        fmt.Println("[Error]: Could not start the server")
        os.Exit(1)
    }
}
