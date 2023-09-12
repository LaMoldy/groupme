package log

import (
	"github.com/LaMoldy/groupme/server/models"
	"github.com/gin-gonic/gin"
)

func CreateLog(c *gin.Context, message string, status int) {
    log := models.Log {
        Message: message,
        Status: status,
    }

    c.IndentedJSON(
        status,
        log,
    )
}
