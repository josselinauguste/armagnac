package query

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/josselinauguste/armagnac/feeds/domain"
)

func TestGetNewItems(t *testing.T) {
	feed := domain.NewFeed("http://lachaineguitare.com/feed/")
	now := time.Now()
	lastSync := now.AddDate(0, 0, -5)
	feed.LastSync = lastSync

	retriever := newFeedGetter(feed)
	newItems, err := retriever.getNewItems()

	assert.Nil(t, err)
	assert.NotEmpty(t, newItems)
	assert.True(t, len(newItems) < 10)
	assert.True(t, feed.LastSync.After(lastSync))
	assert.Equal(t, "La Chaîne Guitare", feed.Title)
}
