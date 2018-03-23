package resource

import (
	"../api"
	"../manager"
	"github.com/gin-gonic/gin"
)

// TweetResource .
type TweetResource struct {
}

// TweetResponse .
type TweetResponse struct {
	FileName    string
	IsStreaming bool
}

// InitRoutes .
func (TweetResource TweetResource) InitRoutes(router *gin.Engine) {
	router.GET("/api/tweets/keyword/:keyword", func(context *gin.Context) {
		keyword := context.Param("keyword")
		if len(keyword) > 0 {
			fileName := api.StartTwitterStreamWithKeywordAndUserID(keyword, "")
			body := TweetResponse{FileName: fileName, IsStreaming: true}
			context.JSON(200, body)
		} else {
			context.String(404, "Not Found")
		}
	})

	router.POST("/api/tweets/stop", func(context *gin.Context) {
		fileName := context.PostForm("fileName")
		if len(fileName) > 0 {
			isStopped := manager.GetChannelManageInstance().StopChannel(fileName)
			body := TweetResponse{FileName: fileName, IsStreaming: !isStopped}
			context.JSON(200, body)
		} else {
			context.String(404, "Not Found")
		}
	})
}
