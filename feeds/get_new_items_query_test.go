package feeds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	feed := newFeed("http://lachaineguitare.com/feed/")
	now := time.Now()
	feed.lastSync = now.AddDate(0, 0, -5)
	currentFeedRepository.Add(feed)
	query := &NewItemsQuery{}

	query.Execute()

	assert.NotEmpty(t, query.NewItems)
	assert.NotEmpty(t, query.NewItems[*feed])
}