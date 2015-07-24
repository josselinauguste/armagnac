package web

import (
	"github.com/josselinauguste/armagnac/feeds"
)

type (
	FeedPresenter struct {
		Feed  feeds.Feed
		Items []feeds.Item
	}

	DigestPresenter struct {
		Feeds []FeedPresenter
	}
)

func newDigestPresenter(query feeds.NewItemsQuery) *DigestPresenter {
	presenter := &DigestPresenter{}
	presenter.Feeds = make([]FeedPresenter, 0, len(query.NewItems))
	for k := range query.NewItems {
		feedPresenter := FeedPresenter{k, query.NewItems[k]}
		presenter.Feeds = append(presenter.Feeds, feedPresenter)
	}

	return presenter
}
