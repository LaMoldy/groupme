package main

import (
    "github.com/gin-gonic/gin"
    "github.com/LaMoldy/groupme/server/pkg/routes"
    "github.com/LaMoldy/groupme/server/pkg/database"
    _ "github.com/lib/pq"
)

func main() {
    // Connect database
    database.ConnectDatabase();

    // Creates the router
    router := gin.Default()

    // Paths
    router.GET("/", routes.Home)

    // Sets up the servers listener
    router.Run("localhost:8080")
}
