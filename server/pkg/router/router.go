package router 

import (
	"github.com/LaMoldy/groupme/server/pkg/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
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

    return router
}
