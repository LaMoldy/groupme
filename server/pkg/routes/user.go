package routes

import (
	"net/http"
	"time"

	"github.com/LaMoldy/groupme/server/models"
	"github.com/LaMoldy/groupme/server/pkg/database"
	"github.com/LaMoldy/groupme/server/pkg/log"
	"github.com/gin-gonic/gin"
)

func GetUserByEmail(c *gin.Context) {
    db := database.ConnectDatabase()

    var users []models.User;

    result := db.Find(&users)
    if result.Error != nil {
	log.CreateLog(
	    c,
	    "A problem has occurred with finding a user",
	    http.StatusInternalServerError,
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
    dob := c.Query("dob")

    dateOfBirth, err := time.Parse("2000-31-6", dob)
    if err != nil {
	log.CreateLog(
	    c,
	    "Error: Problem converting date",
	    http.StatusInternalServerError,
	)
    }

    user := models.User {
	Email: email,
	FirstName: firstName,
	LastName: lastName,
	Password: password,
	DateOfBirth: dateOfBirth,
    }

    result := db.Create(&user)

    if result.Error != nil {
	log.CreateLog(
	    c,
	    "Error: Problem creating user",
	    http.StatusInternalServerError,
	)
    }

    log.CreateLog(
	c,
	"User created successfully",
	http.StatusOK,
    )
}

