package feeds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFeed(t *testing.T) {
	url := "http://salut.com/flux.rss"

	feed := newFeed(url)

	assert.NotNil(t, feed)
	assert.Equal(t, feed.Url, url)
	assert.True(t, feed.lastSync.IsZero())
}
