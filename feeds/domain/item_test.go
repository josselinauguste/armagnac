package domain

import (
	"testing"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/stretchr/testify/assert"
)

func TestNewItemFromRss(t *testing.T) {
	rssItem := rss.Item{
		Title:       "Le titre",
		Description: "La description",
		Author:      rss.Author{Name: "Josselin Auguste"},
		PubDate:     "Fri, 10 Jul 2015 09:33:50 +0000",
		Links:       []*rss.Link{&rss.Link{Href: "http://www.salut.com/le-titre"}},
	}

	item := NewItemFromRss(&rssItem)

	assert.NotNil(t, item)
	assert.Equal(t, rssItem.Title, item.Title)
	assert.Equal(t, rssItem.Description, item.Excerpt)
	pubDate, _ := time.Parse(time.RFC1123Z, rssItem.PubDate)
	assert.Equal(t, pubDate, item.PublicationDate)
	assert.Equal(t, rssItem.Links[0].Href, item.Url)
}

func TestMappingItemWithoutPubDateFailsGracefully(t *testing.T) {
	rssItem := rss.Item{
		Title:       "Le titre",
		Description: "La description",
		Author:      rss.Author{Name: "Josselin Auguste"},
		PubDate:     "LOL, 10 Jul 2015 09:33:50 +0000",
		Links:       []*rss.Link{&rss.Link{Href: "http://www.salut.com/le-titre"}},
	}

	item := NewItemFromRss(&rssItem)

	assert.NotNil(t, item)
	assert.InDelta(t, time.Now().Unix(), item.PublicationDate.Unix(), 10)
}

func TestMappingItemWithoutUrlFailsGracefully(t *testing.T) {
	rssItem := rss.Item{
		Title:       "Le titre",
		Description: "La description",
		Author:      rss.Author{Name: "Josselin Auguste"},
		PubDate:     "Fri, 10 Jul 2015 09:33:50 +0000",
		Links:       []*rss.Link{},
	}

	item := NewItemFromRss(&rssItem)

	assert.NotNil(t, item)
	assert.Empty(t, item.Url)
}
