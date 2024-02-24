package route

import (
	"module/usecase"
	adminfunc "module/utils/admin"

	"github.com/gin-gonic/gin"
)

func RouteAll(r * gin.Engine){
	api := r.Group("api/v1")
	{
		poll := api.Group("/poll")
		{
			poll.POST("/create", usecase.CreatePoll)
			poll.GET("/find",usecase.FindPoll)
			poll.POST("/vote", usecase.Vote)
			poll.GET("/stat", usecase.ShowStat)
		}
		block := api.Group("block")
		{
			block.GET("/find", usecase.FindBlock)
		}

	}
	admin := r.Group("/admin")
	{
		admin.GET("/dropalltables",adminfunc.DropAllTables)
	}
}

