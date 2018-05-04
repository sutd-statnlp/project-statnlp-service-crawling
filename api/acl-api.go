package api

import (
	"strings"

	"github.com/gocolly/colly"
)

const (
	url = "http://acl2018.org/conference/accepted-papers/"
)

// StartCrawlACLAuthorsAccepted .
func StartCrawlACLAuthorsAccepted() []string {
	collector := colly.NewCollector()
	var authors []string
	collector.OnHTML(".listing", func(e *colly.HTMLElement) {
		for _, author := range GetAuthorsFromString(e.ChildText(".paper-authors")) {
			authors = append(authors, author)
		}
	})
	collector.Visit(url)
	return RemoveDuplicateInSlice(authors)
}

// StartCrawlACLLastAuthorsAccepted .
func StartCrawlACLLastAuthorsAccepted() []string {
	collector := colly.NewCollector()
	var authors []string
	collector.OnHTML(".listing", func(e *colly.HTMLElement) {
		author := GetLastAuthor(e.ChildText(".paper-authors"))
		authors = append(authors, author)
	})
	collector.Visit(url)
	return RemoveDuplicateInSlice(authors)
}

// StartCrawlACLLastUniqueAuthorsAccepted .
func StartCrawlACLLastUniqueAuthorsAccepted() []string {
	collector := colly.NewCollector()
	var authors []string
	collector.OnHTML(".listing", func(e *colly.HTMLElement) {
		author := GetLastUniqueAuthor(e.ChildText(".paper-authors"))
		if len(author) > 0 {
			authors = append(authors, author)
		}
	})
	collector.Visit(url)
	return RemoveDuplicateInSlice(authors)
}

// GetLastAuthor .
func GetLastAuthor(authors string) string {
	var result string
	splitString := strings.Split(authors, "and")
	if len(splitString) > 1 {
		result = splitString[1]
	} else {
		result = splitString[0]
	}
	return strings.TrimSpace(result)
}

// GetLastUniqueAuthor .
func GetLastUniqueAuthor(authors string) string {
	splitString := strings.Split(authors, "and")
	if len(splitString) == 1 {
		return splitString[0]
	}
	return ""
}

// GetAuthorsFromString .
func GetAuthorsFromString(authorsRow string) []string {
	var result []string
	result = strings.Split(authorsRow, ", ")
	splitString := strings.Split(result[len(result)-1], "and")
	result = result[:len(result)-1]
	for _, author := range splitString {
		result = append(result, strings.TrimSpace(author))
	}
	return result
}

// RemoveDuplicateInSlice .
func RemoveDuplicateInSlice(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
