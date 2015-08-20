package web

import (
	"testing"

	"github.com/josselinauguste/armagnac/feeds/domain"
	"github.com/josselinauguste/armagnac/feeds/query"
	"github.com/stretchr/testify/assert"
)

func TestNewDigestPresenter(t *testing.T) {
	query := query.NewNewItemsQuery()
	query.NewItems = make(map[domain.Feed][]domain.Item)
	feed := domain.NewFeed("http://salut.com")
	items := make([]domain.Item, 1)
	query.NewItems[*feed] = items

	presenter := newDigestPresenter(*query)

	assert.NotNil(t, presenter)
	assert.Len(t, presenter.Feeds, 1)
	assert.NotNil(t, presenter.Feeds[0].Feed)
	assert.Len(t, presenter.Feeds[0].Items, 1)
}
