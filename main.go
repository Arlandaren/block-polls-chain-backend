package main

import (
	"module/models"
	"module/models/blockchain"
	"module/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main(){
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        defer file.Close()

        logger.SetOutput(file)
    } else {

        logrus.Warn("Не удалось создать файл логов: ", err)
    }

	logger.Info("Starting")
	gin.DefaultWriter = logger.Writer()
	r := gin.Default()

	logger.Info("Loading env variables")
	err = godotenv.Load()
	if err != nil{
		logger.Fatalf("Couldnt load env variables")
	}

	Config := models.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLmode: os.Getenv("DB_SSL"),
		DBname: os.Getenv("DB_NAME"),
	}

	logger.Info("Initializing database")
	models.InitDB(Config)
	logger.Info("Postgres succesfully initialized")

	blockchain.InitDB()
	logger.Info("Mongo succesfully initialized")

	route.RouteAll(r)
	logger.Info("Router ready to route")

	logger.Info("Server listening on 8080")
	r.Run(":8080")

}