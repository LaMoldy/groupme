package main

import (
    "github.com/LaMoldy/groupme/server/pkg/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Creates the router
    router := gin.Default()

    // Paths
    router.GET("/", routes.Home)

    // Sets up the servers listener
    router.Run("localhost:8080")
}
