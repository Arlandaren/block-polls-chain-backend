package adminfunc

import (
	"module/models"

	"github.com/gin-gonic/gin"
)

func DropAllTables(c *gin.Context){
	models.DropAllTables()
	c.JSON(200,gin.H{"Message":"succesfully deleted"})
}