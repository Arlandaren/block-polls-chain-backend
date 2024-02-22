package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

var DB *gorm.DB

func InitDB(cfg Config){
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",cfg.Host,cfg.Port,cfg.User,cfg.Password,cfg.DBname,cfg.SSLmode)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err.Error())
	}
	// if err := con.AutoMigrate(&User{},&Blockchain{},&Poll{},&Option{},&Vote{}); err!=nil{
	// 	panic(err)
	// }
	if err := con.AutoMigrate(&User{},&Poll{},&Option{},&Vote{}); err!=nil{
		panic(err)
	}
	DB = con
}

func CreatePoll(block string,title string){
	var poll Poll
	poll.Block = block
	poll.Title = title
	DB.Create(&poll)
}
func CreateOption(block string, text string, poll_block string){
	var option Option
	option.Block = block
	option.Text = text
	option.PollBlock = poll_block

	if err := DB.Create(&option).Error; err != nil {
        fmt.Println("error creating option:",err)
    }
}

func FindPoll(hash string) (*Poll,*[]Option,error){
	var poll Poll
	err := DB.Where("block=?",hash).First(&poll).Error
	if err != nil{
		return nil,nil,err
	}
	var options []Option
	err = DB.Where("poll_block=?",poll.Block).Find(&options).Error
	if err != nil{
		return nil,nil,err
	}
	return &poll,&options,nil

}

func DropAllTables() error {
	if err := DB.Migrator().DropTable(&User{}, &Poll{}, &Option{}, &Vote{}); err != nil {
        return err
    }
	return nil
}