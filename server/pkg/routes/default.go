package routes

import (
	"net/http"

	"github.com/LaMoldy/groupme/server/pkg/log"
	"github.com/gin-gonic/gin"
)

// Sets up default route handler
func Welcome(c *gin.Context) {
    log.CreateLog(
	c,
	"Welcome to the API!",
	http.StatusOK,
    )
}

func RedirectToWelcome(c *gin.Context) {
    c.Redirect(http.StatusFound, "/api")
}

