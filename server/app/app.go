package app

/*
	Main entry point for the server. Instantiates the main app, setups the gin http server on port 8080
	Handles CORS hearders

	call app.Init() to start the server

	sets up the app log found in logs/app.log
*/

import (
	"github.com/gin-gonic/gin"
	"rundoo.com/server/apis"
	"log"
	"os"
	"rundoo.com/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}

			c.Next()
	}
}

func Init() {
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("API server running")
	router := gin.Default() 
	router.SetTrustedProxies(nil)
	router.Use(CORSMiddleware())
	router.POST("/product", apis.CreateProduct)
	router.POST("/searchProducts", apis.SearchProducts)
	router.Run(config.GetServerURL())
}
