package feeds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewItems(t *testing.T) {
	feed := newFeed("http://lachaineguitare.com/feed/")
	now := time.Now()
	lastSync := now.AddDate(0, 0, -5)
	feed.lastSync = lastSync

	retriever := newFeedGetter(feed)
	newItems, err := retriever.getNewItems()

	assert.Nil(t, err)
	assert.NotEmpty(t, newItems)
	assert.True(t, len(newItems) < 10)
	assert.True(t, feed.lastSync.After(lastSync))
}
