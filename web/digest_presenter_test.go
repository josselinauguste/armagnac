package web

import (
	"testing"

	"github.com/josselinauguste/armagnac/feeds/domain"
	"github.com/josselinauguste/armagnac/feeds/query"
	"github.com/stretchr/testify/assert"
)

func TestNewDigestPresenter(t *testing.T) {
	feed := domain.NewFeed("http://salut.com")
	feed.ID = "id"
	query := query.NewNewItemsQuery()
	query.Feeds = []*domain.Feed{feed}
	items := make([]domain.Item, 1)
	query.NewItems[feed.ID] = items

	presenter := newDigestPresenter(*query)

	assert.NotNil(t, presenter)
	assert.Len(t, presenter.Feeds, 1)
	assert.NotNil(t, presenter.Feeds[0].Feed)
	assert.Len(t, presenter.Feeds[0].Items, 1)
}
