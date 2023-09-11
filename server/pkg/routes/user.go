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

func CreateUser(c *gin.Context) {
    db := database.ConnectDatabase()

    email := c.Query("email")
    firstName := c.Query("firstName")
    lastName := c.Query("lastName")
    password := c.Query("password")

    user := models.User {
	Email: email,
	FirstName: firstName,
	LastName: lastName,
	Password: password,
    }

    result := db.Create(&user)

    if result.Error != nil {
	var message = models.Message {
	    Message: "Error: Problem creating a user",
	    Status: http.StatusInternalServerError,
	}

	c.IndentedJSON(
	    http.StatusInternalServerError,
	    message,
	)
    }

    var message = models.Message {
	Message: "User created successfully",
	Status: http.StatusOK,
    }

    c.IndentedJSON(
	http.StatusOK,
	message,
    )
}
