package main

import (
	"github.com/gin-gonic/gin"

	"./controller"
	"./resource"
)

func main() {
	router := setupRoutes()

	router.Run(":8220")
}

func setupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Next()
	})

	homeController := controller.HomeController{}
	homeController.InitRoutes(router)

	tweetResource := resource.TweetResource{}
	tweetResource.InitRoutes(router)

	return router
}
