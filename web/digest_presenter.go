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
	presenter.Feeds = make([]FeedPresenter, 0, len(query.NewItems))
	for k := range query.NewItems {
		feedPresenter := FeedPresenter{k, query.NewItems[k]}
		presenter.Feeds = append(presenter.Feeds, feedPresenter)
	}

	return presenter
}
