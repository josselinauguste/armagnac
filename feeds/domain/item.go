package domain

import (
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
)

type Item struct {
	Title           string
	Url             string
	Description     string
	PublicationDate time.Time
}

func NewItemFromRss(rssItem *rss.Item) *Item {
	description := extractDescription(rssItem)
	item := &Item{Title: rssItem.Title, Description: description}
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

func extractDescription(rssItem *rss.Item) string {
	if len(rssItem.Description) > 0 {
		return rssItem.Description
	}
	if rssItem.Content != nil {
		return rssItem.Content.Text
	}
	return ""
}

func (item Item) Excerpt() string {
	return item.Description
}
