package feeds

import (
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
)

type Item struct {
	Title           string
	Url             string
	Excerpt         string
	PublicationDate time.Time
}

func newItemFromRss(rssItem *rss.Item) *Item {
	url := rssItem.Links[0].Href
	publicationDate, err := rssItem.ParsedPubDate()
	if err != nil {
		publicationDate = time.Now()
	}
	return &Item{Title: rssItem.Title, Url: url, Excerpt: rssItem.Description, PublicationDate: publicationDate}
}
