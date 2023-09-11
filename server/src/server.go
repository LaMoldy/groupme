package main

import (
    "fmt"
    "net/http"
    "time"
    "os"

    "github.com/LaMoldy/groupme/server/pkg/routes"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    // Creates the router
    router := gin.Default()

    // Router configuration
    router.RedirectFixedPath = true
    router.RedirectTrailingSlash = true

    // Paths
    router.GET("/", routes.RedirectToWelcome)
    router.GET("/api", routes.Welcome)
    router.GET("/api/user/email", routes.GetUserByEmail)
    router.POST("/api/user/create", routes.CreateUser)

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
