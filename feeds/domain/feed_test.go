package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFeed(t *testing.T) {
	url := "http://salut.com/flux.rss"

	feed := NewFeed(url)

	assert.NotNil(t, feed)
	assert.Equal(t, feed.Uri, url)
	assert.True(t, feed.LastSync.IsZero())
}
