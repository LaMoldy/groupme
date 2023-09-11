package routes

import (
    "net/http"

    "github.com/LaMoldy/groupme/server/models"
    "github.com/gin-gonic/gin"
)

// Sets up default route handler
func Welcome(c *gin.Context) {
    var message = models.Message {
	Message: "Welcome to GroupMe's API!",
	Status: http.StatusOK,
    }

    c.IndentedJSON(
	http.StatusOK,
	message,
    )
}

func RedirectToWelcome(c *gin.Context) {
    c.Redirect(http.StatusFound, "/api")
}

