package query

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/josselinauguste/armagnac/feeds/domain"
	"github.com/josselinauguste/armagnac/feeds/repository"
)

func TestExecute(t *testing.T) {
	feed := domain.NewFeed("http://lachaineguitare.com/feed/")
	now := time.Now()
	feed.LastSync = now.AddDate(0, 0, -5)
	repository.CurrentFeedRepository.Persist(feed)
	query := &NewItemsQuery{}

	query.Execute()

	assert.NotEmpty(t, query.NewItems)
	assert.NotEmpty(t, query.NewItems[*feed])
}
