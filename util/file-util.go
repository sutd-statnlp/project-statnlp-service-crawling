package util

import (
	"encoding/csv"
	"log"
	"os"
)

const tweetResultFileName = "tweets.csv"
const dataDirectory = "./data/"

var tweetCSVColumns = []string{"Timestamp",
	"ID", "Text", "UserName", "CreatedAt",
	"FavoriteCount", "Favorited", "FilterLevel", "Lang",
	"PlaceName", "RetweetCount", "Retweeted"}

// InitTweetFile .
func InitTweetFile() string {
	fileName := GetCurrentTimestamp() + "-" + tweetResultFileName
	saveToCsv(tweetCSVColumns, true, fileName)
	return fileName
}

// SaveTweetResult .
func SaveTweetResult(data []string, fileName string) {
	saveToCsv(data, false, fileName)
}

// SaveToCsv : Save data into CSV file.
func saveToCsv(data []string, attempt bool, fileName string) {
	var file *os.File
	var err error
	if attempt {
		file, err = os.Create(dataDirectory + fileName)
	} else {
		file, err = os.OpenFile(dataDirectory+fileName, os.O_WRONLY|os.O_APPEND, 0644)
	}
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	checkError("Cannot write to file", writer.Write(data))
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
