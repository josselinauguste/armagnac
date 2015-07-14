package feeds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewItems(t *testing.T) {
	feed := NewFeed("http://lachaineguitare.com/feed/")
	now := time.Now()
	feed.lastSync = now.AddDate(0, 0, -5)

	retriever := NewFeedGetter(feed)
	newItems, err := retriever.RetrieveNewItems()

	assert.Nil(t, err)
	assert.NotEmpty(t, newItems)
	assert.True(t, len(newItems) < 10)
}
