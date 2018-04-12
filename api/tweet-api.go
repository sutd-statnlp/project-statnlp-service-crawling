package api

import (
	"strconv"
	"strings"
	"time"

	"../config"
	"../manager"
	"../util"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func getClient() *twitter.Client {
	oAuthConfig := oauth1.NewConfig(config.TwitteConsumeKey, config.TwitteConsumeSecret)
	token := oauth1.NewToken(config.TwitterAccessToken, config.TwitteAccessSecret)
	httpClient := oAuthConfig.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

// StartTwitterStream .
func StartTwitterStream(keywords string, locations string, userIds string, languages string, maxMinute int) string {
	fileName := util.InitTweetFile()
	channel := stream(fileName, keywords, locations, userIds, languages, false)
	manager.GetChannelManageInstance().AddChannel(fileName, channel)
	go func() {
		time.Sleep(time.Duration(maxMinute+1) * time.Minute)
		manager.GetChannelManageInstance().StopChannel(fileName)
	}()
	return fileName
}

// StartTwitterStream .
func stream(fileName string, keywords string, locations string, userIds string, languages string, stallWarning bool) chan string {
	channel := make(chan string)
	go func() {
		client := getClient()
		params := &twitter.StreamFilterParams{
			Track:         strings.Split(keywords, ","),
			StallWarnings: twitter.Bool(stallWarning),
			Locations:     strings.Split(locations, ","),
			Language:      strings.Split(languages, ","),
			Follow:        strings.Split(userIds, ","),
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
