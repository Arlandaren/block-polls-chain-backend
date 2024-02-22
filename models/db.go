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
	if err := con.AutoMigrate(&User{},&Blockchain{},&Poll{},&Option{},&Vote{}); err!=nil{
		panic(err)
	}
	DB = con
}