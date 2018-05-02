package api

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

// StartCrawlACLAuthorsAccepted .
func StartCrawlACLAuthorsAccepted(context *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("acl2018.org"),
		colly.Async(true),
	)

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		context.Stream(func(w io.Writer) bool {
			context.SSEvent("author", e.Name)
			return true
		})
	})

	collector.Visit("http://acl2018.org/")
}
