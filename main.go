package main

import (
	"fmt"
	"github.com/sandeepmendiratta/gin-go-boilerplate/app"
	log "github.com/sirupsen/logrus"
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	// Will log all level down from debug(highest to lowest)
	log.SetLevel(log.DebugLevel)

	fmt.Println("Welcome to the app")
	app.StartApp()
}
