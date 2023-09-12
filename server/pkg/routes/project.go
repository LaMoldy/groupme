package routes

import (
	"net/http"

	"github.com/LaMoldy/groupme/server/models"
	"github.com/LaMoldy/groupme/server/pkg/database"
	"github.com/LaMoldy/groupme/server/pkg/log"
	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {

}

func UpdateProject(c *gin.Context) {

}

func DeleteProject(c *gin.Context) {
    db := database.ConnectDatabase()

    projectId := c.Query("projectId")

    db.Delete(models.Project{}, projectId)

    log.CreateLog(
	c,
	"Successfully deleted the project",
	http.StatusOK,
    )
}
