package route

import (
	"module/usecase"

	"github.com/gin-gonic/gin"
)

func RouteAll(r * gin.Engine){
	api := r.Group("api/v1")
	{
		api.POST("/create", usecase.CreatePoll)
	}
}