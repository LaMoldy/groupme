package routes

import (
	"net/http"

	"github.com/LaMoldy/groupme/server/models"
	"github.com/LaMoldy/groupme/server/pkg/database"
	"github.com/gin-gonic/gin"
)

func GetUserByEmail(c *gin.Context) {
    db := database.ConnectDatabase()

    var users []models.User;

    result := db.Find(&users)
    if result.Error != nil {
	c.IndentedJSON(
	    http.StatusInternalServerError,
	    models.Message {
		Message: "A problem has occurred: " + result.Error.Error(),
		Status: http.StatusInternalServerError,
	    },
	)
	return
    }

    var foundUser models.User
    email := c.Query("email")

    for _, user := range users {
	if user.Email == email {
	    foundUser = user
	    break
	}
    }

    c.IndentedJSON(
	http.StatusOK,
	foundUser,
    )
}
