package resource

import (
	"../api"
	"github.com/gin-gonic/gin"
)

// InitAclRoutes .
func InitAclRoutes(router *gin.Engine) {
	router.GET("/api/acl/authors/accepted", func(context *gin.Context) {
		authorsRows := api.StartCrawlACLAuthorsAccepted()
		context.JSON(200, authorsRows)
	})
	router.GET("/api/acl/authors/accepted/last", func(context *gin.Context) {
		authorsRows := api.StartCrawlACLLastAuthorsAccepted()
		context.Header("Content-Type", "application/json; charset=utf-8")
		context.JSON(200, authorsRows)
	})
	router.GET("/api/acl/authors/accepted/last/unique", func(context *gin.Context) {
		authorsRows := api.StartCrawlACLLastUniqueAuthorsAccepted()
		context.Header("Content-Type", "application/json; charset=utf-8")
		context.JSON(200, authorsRows)
	})
}
