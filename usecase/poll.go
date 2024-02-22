package usecase

import (
	"module/models/blockchain"

	"github.com/gin-gonic/gin"
)

func CreatePoll(c *gin.Context){
	data := c.Query()
	block := blockchain.AddBlock("dad",123)
	c.JSON(200,gin.H{"message":"Success, block added","block":block})
}