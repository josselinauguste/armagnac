package web

import (
	"github.com/josselinauguste/armagnac/feeds/domain"
	"github.com/josselinauguste/armagnac/feeds/query"
)

type (
	FeedPresenter struct {
		Feed  domain.Feed
		Items []domain.Item
	}

	DigestPresenter struct {
		Feeds []FeedPresenter
	}
)

func newDigestPresenter(query query.NewItemsQuery) *DigestPresenter {
	presenter := &DigestPresenter{}
	presenter.Feeds = make([]FeedPresenter, 0, len(query.Feeds))
	for _, feed := range query.Feeds {
		feedPresenter := FeedPresenter{*feed, query.NewItems[feed.ID]}
		presenter.Feeds = append(presenter.Feeds, feedPresenter)
	}

	return presenter
}
