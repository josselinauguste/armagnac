package query

import (
	rss "github.com/jteeuwen/go-pkg-rss"

	"github.com/josselinauguste/armagnac/feeds/domain"
	"github.com/josselinauguste/armagnac/feeds/repository"
)

type feedGetter struct {
	feed          *domain.Feed
	feedConnector *rss.Feed
	items         []*rss.Item
}

func newFeedGetter(feed *domain.Feed) *feedGetter {
	getter := new(feedGetter)
	getter.feed = feed
	getter.feedConnector = rss.NewWithHandlers(5, true, getter, getter)
	return getter
}

func (getter *feedGetter) getNewItems() ([]domain.Item, error) {
	if err := getter.feedConnector.Fetch(getter.feed.Uri, nil); err != nil {
		return nil, err
	}
	newItems := make([]domain.Item, 0, len(getter.items))
	for _, rssItem := range getter.items {
		if pubDate, _ := rssItem.ParsedPubDate(); pubDate.After(getter.feed.LastSync) {
			item := domain.NewItemFromRss(rssItem)
			newItems = append(newItems, *item)
			if pubDate.After(getter.feed.LastSync) {
				getter.feed.LastSync = pubDate
			}
		}
	}
	repository.CurrentFeedRepository.Persist(getter.feed)
	return newItems, nil
}

func (getter *feedGetter) ProcessChannels(feed *rss.Feed, newchannels []*rss.Channel) {
	getter.feed.Title = getter.feedConnector.Channels[0].Title
}

func (getter *feedGetter) ProcessItems(f *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	getter.items = append(getter.items, newitems...)
}
