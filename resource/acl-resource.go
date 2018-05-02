package resource

import (
	"../api"
	"github.com/gin-gonic/gin"
)

// InitAclRoutes .
func InitAclRoutes(router *gin.Engine) {
	router.GET("/api/acl/authors/accepted", func(context *gin.Context) {
		api.StartCrawlACLAuthorsAccepted(context)
	})
}
