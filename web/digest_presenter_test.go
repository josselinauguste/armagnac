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
	feed.Title = "Salut"
	query := query.NewNewItemsQuery()
	query.Feeds = []*domain.Feed{feed}
	items := make([]domain.Item, 1)
	query.NewItems[feed.ID] = items

	presenter := newDigestPresenter(*query)

	assert.NotNil(t, presenter)
	assert.Len(t, presenter.Feeds, 1)
	assert.NotEmpty(t, presenter.Feeds[0].Title)
	assert.Len(t, presenter.Feeds[0].Entries, 1)
}

func TestUnescapeExcerpt(t *testing.T) {
	presenter := EntryPresenter{"", "&lt;15", ""}

	assert.Equal(t, "<15", string(presenter.FormattedExcerpt()))
}

func TestSanitizeExcerpt(t *testing.T) {
	presenter := EntryPresenter{"", "&lt;a href=\"javascript:alert('XSS1')\" onmouseover=\"alert('XSS2')\"&gt;XSS&lt;/a&gt;", ""}

	assert.Equal(t, "XSS", string(presenter.FormattedExcerpt()))
}
