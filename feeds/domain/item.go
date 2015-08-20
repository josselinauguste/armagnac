package domain

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

func NewItemFromRss(rssItem *rss.Item) *Item {
	item := &Item{Title: rssItem.Title, Excerpt: rssItem.Description}
	if len(rssItem.Links) > 0 {
		item.Url = rssItem.Links[0].Href
	}
	if publicationDate, err := rssItem.ParsedPubDate(); err != nil {
		item.PublicationDate = time.Now()
	} else {
		item.PublicationDate = publicationDate
	}
	return item
}
