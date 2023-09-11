package main

import (
    "github.com/gin-gonic/gin"
    "github.com/LaMoldy/groupme/server/pkg/routes"
    _ "github.com/lib/pq"
)

func main() {
    // Creates the router
    router := gin.Default()

    // Paths
    router.GET("/", routes.Home)
    router.GET("/user", routes.GetUserByEmail)

    // Sets up the servers listener
    router.Run("localhost:8080")
}
