package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)



func init() {
	//loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}
var router *gin.Engine
func StartApp() {
	router = gin.Default()
	initializeRoutes()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)

	}
}