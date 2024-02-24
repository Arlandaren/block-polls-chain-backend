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

type vote struct{
	Poll_block string `json:"poll_block"`
	Option_block string `json:"option_block"`
}

func CreatePoll(c *gin.Context){
	var poll poll
	if err := c.ShouldBindJSON(&poll); err != nil{
		c.JSON(400,gin.H{"message":"incorrect body of post method request","status":"error"})
		return
	}
	poll_block, err := blockchain.AddBlock(fmt.Sprintf("Poll:%s",poll.Title))
	if err != nil{
		c.JSON(500,gin.H{"message":"couldnt create block in blockchain", "status":"error","error":err})
		return
	}
	if err:=models.CreatePoll(poll_block.Hash,poll.Title);err!=nil{
		c.JSON(500,gin.H{"message":"couldnt insert poll in db", "status":"error","error":err})
		return
	}
	for i,v:= range poll.Options{
		option_block, err := blockchain.AddBlock(fmt.Sprintf("Option:%s",v))
		if err != nil{
			c.JSON(500,gin.H{"message":"couldnt create block in blockchain", "status":"error","error":err})
			return
		}
		if err:=models.CreateOption(option_block.Hash,fmt.Sprintf("Options:%d.%s",i+1,v),poll_block.Hash);err!=nil{
			c.JSON(500,gin.H{"message":"couldnt insert poll in db", "status":"error","error":err})
			return
		}
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

func Vote(c *gin.Context){
	var vote vote
	if err := c.ShouldBindJSON(&vote); err != nil{
		c.JSON(400,gin.H{"message":"incorrect body of post method request","status":"error"})
		return
	}
	block,err := blockchain.AddBlock(fmt.Sprintf("Vote:option:%s, poll:%s",vote.Option_block,vote.Poll_block))
	if err != nil{
		c.JSON(500,gin.H{"message":"couldnt create block in blockchain", "status":"error","error":err})
		return
	}
	if err:= models.CreateVote(block.Hash,vote.Poll_block,vote.Option_block); err!=nil{
		c.JSON(500,gin.H{"message":"couldnt insert vote in db", "status":"error","error":err})
		return
	}
	c.JSON(200,gin.H{"message":"vote created","status":"success"})
	
}
func ShowStat(c *gin.Context){
	poll_block := c.Query("poll")
	option_block := c.Query("option")
	count := models.Stats(poll_block,option_block)
	c.JSON(200,gin.H{"count":count,"status":"success"})
}
