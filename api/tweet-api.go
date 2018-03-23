package api

import (
	"fmt"
	"strconv"

	"../manager"
	"../util"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const accessToken string = "830802762753417218-YqYSKPmmn1wVvKItENWPhHF8f1uDjcJ"
const accessSecret string = "2SJLtyeHRqc23jtEnOKM3eWcTbL9CUHjsncb4CgCpR1HR"
const consumeKey string = "YUXDDDmE0AvI8e0diUdsbJ7ph"
const consumeSecret string = "vNbYpJjJzy5ZdqVZlqqk1MKlspviQOIj1Txnuw69rUfbClkdy5"

func getClient() *twitter.Client {
	config := oauth1.NewConfig(consumeKey, consumeSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

// InitTwitterStream : test streaming.
func InitTwitterStream(keywords []string, userIds []string, stallWarning bool) string {
	fileName := util.InitTweetFile()
	channel := StartTwitterStream(fileName, keywords, userIds, stallWarning)
	<-channel
	fmt.Println("Done channel ", fileName)
	return fileName
}

// StartTwitterStreamWithKeywordAndUserID : start streaming.
func StartTwitterStreamWithKeywordAndUserID(keyword string, userID string) string {
	fileName := util.InitTweetFile()
	channel := StartTwitterStream(fileName, []string{keyword}, []string{userID}, false)
	manager.GetChannelManageInstance().ChannelMap[fileName] = channel
	return fileName
}

// StartTwitterStream .
func StartTwitterStream(fileName string, keywords []string, userIds []string, stallWarning bool) chan string {
	channel := make(chan string)
	go func() {
		client := getClient()
		params := &twitter.StreamFilterParams{
			Track:         keywords,
			StallWarnings: twitter.Bool(stallWarning),
			Follow:        userIds,
		}

		demux := twitter.NewSwitchDemux()
		demux.Tweet = func(tweet *twitter.Tweet) {
			saveTweet(fileName, tweet)
		}
		stream, _ := client.Streams.Filter(params)
		go demux.HandleChan(stream.Messages)

		channel <- fileName
		stream.Stop()

	}()
	return channel
}

func saveTweet(fileName string, tweet *twitter.Tweet) {
	timestamp := util.GetCurrentTimestamp()
	fmt.Println("Timestamp ", timestamp, " of ", tweet.ID)
	util.SaveTweetResult([]string{
		timestamp,
		tweet.IDStr,
		tweet.Text,
		tweet.User.Name,
		tweet.CreatedAt,
		strconv.Itoa(tweet.FavoriteCount),
		strconv.FormatBool(tweet.Favorited),
		tweet.FilterLevel,
		tweet.Lang,
		getPlaceName(tweet),
		strconv.Itoa(tweet.RetweetCount),
		strconv.FormatBool(tweet.Retweeted),
	}, fileName)
}

func getPlaceName(tweet *twitter.Tweet) string {
	if tweet.Place != nil {
		return tweet.Place.Name
	}
	return "none"
}
