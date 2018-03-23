package controller

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// HomeController .
type HomeController struct {
}

// InitRoutes .
func (homeController HomeController) InitRoutes(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./view", true)))
	router.Use(static.Serve("/data", static.LocalFile("./data", true)))
}
