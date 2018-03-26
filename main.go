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

	controller.InitHomeRoutes(router)
	resource.InitTweetRoutes(router)

	return router
}
