package web

import (
	"html"
	"html/template"

	"github.com/josselinauguste/armagnac/feeds/query"
	"github.com/microcosm-cc/bluemonday"
)

type (
	EntryPresenter struct {
		Title   string
		Excerpt string
		Url     string
	}

	FeedPresenter struct {
		Title   string
		Entries []EntryPresenter
	}

	DigestPresenter struct {
		Feeds []FeedPresenter
	}
)

func newDigestPresenter(query query.NewItemsQuery) *DigestPresenter {
	digestPresenter := &DigestPresenter{}
	digestPresenter.Feeds = make([]FeedPresenter, 0, len(query.Feeds))
	for _, feed := range query.Feeds {
		feedPresenter := FeedPresenter{feed.Title, make([]EntryPresenter, 0, len(query.NewItems[feed.ID]))}
		for _, item := range query.NewItems[feed.ID] {
			entryPresenter := EntryPresenter{item.Title, item.Excerpt(), item.Url}
			feedPresenter.Entries = append(feedPresenter.Entries, entryPresenter)
		}
		digestPresenter.Feeds = append(digestPresenter.Feeds, feedPresenter)
	}
	return digestPresenter
}

func (presenter EntryPresenter) FormattedExcerpt() template.HTML {
	unescaped := html.UnescapeString(presenter.Excerpt)
	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(unescaped)
	unescaped = html.UnescapeString(sanitized)
	return template.HTML(unescaped)
}
