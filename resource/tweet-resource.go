package resource

import (
	"../api"
	"../manager"
	"github.com/gin-gonic/gin"
)

// TweetResponse .
type TweetResponse struct {
	FileName    string
	IsStreaming bool
}

// InitTweetRoutes .
func InitTweetRoutes(router *gin.Engine) {

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

	router.POST("/api/tweets/stream", func(context *gin.Context) {
		keyword := context.PostForm("keyword")
		location := context.PostForm("location")
		language := context.PostForm("language")
		userID := context.PostForm("userId")
		fileName := api.StartTwitterStream(keyword, location, userID, language)
		body := TweetResponse{FileName: fileName, IsStreaming: true}
		context.JSON(200, body)
	})
}
