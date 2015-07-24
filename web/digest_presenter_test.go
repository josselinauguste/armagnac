package web

import (
	"testing"

	"github.com/josselinauguste/armagnac/feeds"
	"github.com/stretchr/testify/assert"
)

func TestNewDigestPresenter(t *testing.T) {
	query := feeds.NewNewItemsQuery()
	query.NewItems = make(map[feeds.Feed][]feeds.Item)
	feed := feeds.NewFeed("http://salut.com")
	items := make([]feeds.Item, 1)
	query.NewItems[*feed] = items

	presenter := newDigestPresenter(*query)

	assert.NotNil(t, presenter)
	assert.Len(t, presenter.Feeds, 1)
	assert.NotNil(t, presenter.Feeds[0].Feed)
	assert.Len(t, presenter.Feeds[0].Items, 1)
}
