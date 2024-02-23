package usecase

import (
	"module/models/blockchain"

	"github.com/gin-gonic/gin"
)

func FindBlock(c *gin.Context){
	hash := c.Query("hash")
	block, err:= blockchain.FindByHash(hash)
	if err != nil{
		c.JSON(500,gin.H{"message":"couldnt find block by hash","status":"error"})
		return
	}
	c.JSON(500,gin.H{"block":block,"status":"success"})
}