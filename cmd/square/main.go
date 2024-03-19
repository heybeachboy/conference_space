package main

import (
	"ConferenceSpace/bootstrap"
	_ "ConferenceSpace/cmd/square/docs"
	route2 "ConferenceSpace/route"
	//docs "github.com/go-project-name/docs"
	"flag"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

var configFilePath string

func init() {
	flag.Parse()
	flag.StringVar(&configFilePath, "c", "./config/config.yaml", "config file path")
}

// @title Reference Square API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8900
// @BasePath /conference
func main() {
	if err := bootstrap.BootInit(configFilePath); err != nil {
		log.Printf("conference space boot init error : %s", err.Error())
		return
	}
	gin := route2.RegisterRoute()
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := gin.Run(":8900"); err != nil {
		log.Printf("run http server error : %s", err.Error())
	}
}
