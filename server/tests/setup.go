package main

import (
	"github.com/gin-gonic/gin"
	//"server/app"
	//"server/config"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func loadConfig() {
	//config.AppConfig = config.GetAppConfig()
}

func setup() {
	// appConfig := config.GetAppConfig()
	// isDev := false
	// if appConfig.IsDevelopment() {
	// 	isDev = true
	// }
	// isTest := false
	// if appConfig.IsTesting() {
	// 	isTest = true
	// }

	// app.SetupDatabase(isDev, isTest)
}

func tearDown() {

}