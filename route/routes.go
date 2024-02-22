package route

import (
	"module/usecase"
	adminfunc "module/utils/admin"

	"github.com/gin-gonic/gin"
)

func RouteAll(r * gin.Engine){
	api := r.Group("api/v1")
	{
		api.POST("/create", usecase.CreatePoll)
		api.GET("/find_poll",usecase.FindPoll)
	}
	admin := r.Group("/admin")
	{
		admin.GET("/dropalltables",adminfunc.DropAllTables)
	}
}