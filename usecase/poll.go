package usecase

import (
	"fmt"
	"module/models"
	"module/models/blockchain"

	"github.com/gin-gonic/gin"
)

type poll struct{
	Title string `json:"Title"`
	Options []string `json:"Options"`
	Owner uint64 `json:"Owner"`
}

func CreatePoll(c *gin.Context){
	var poll poll
	if err := c.ShouldBindJSON(&poll); err != nil{
		c.JSON(400,gin.H{"message":"incorrect body of post method request","status":"error"})
	}
	poll_block, err := blockchain.AddBlock(fmt.Sprintf("Poll:%s",poll.Title),poll.Owner)
	if err != nil{
		c.JSON(500,gin.H{"message":"couldnt create block in blockchain", "status":"error","error":err})
	}
	models.CreatePoll(poll_block.Hash,poll.Title)
	for i,v:= range poll.Options{
		option_block, err := blockchain.AddBlock(fmt.Sprintf("Option:%s",v),poll.Owner)
		if err != nil{
		c.JSON(500,gin.H{"message":"couldnt create block in blockchain", "status":"error","error":err})
		}
		models.CreateOption(option_block.Hash,fmt.Sprintf("Options:%d.%s",i+1,v),poll_block.Hash)
	}
	c.JSON(200,gin.H{"message":"Poll created","status":"success","hash":poll_block.Hash})
}
func FindPoll(c *gin.Context){
	hash := c.Query("hash")
	poll,options,err := models.FindPoll(hash)
	if err != nil{
		c.JSON(500,gin.H{"message":fmt.Sprintf("couldnt find poll with hash: %s",hash),"status":"error"})
		return
	}
	c.JSON(200,gin.H{"status":"success","poll":gin.H{"options":options,"title":poll.Title,"hash":poll.Block}})
}
