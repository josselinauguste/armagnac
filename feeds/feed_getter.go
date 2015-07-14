package feeds

import rss "github.com/jteeuwen/go-pkg-rss"

type feedGetter struct {
	feed          *Feed
	feedConnector *rss.Feed
	items         []*rss.Item
}

func newFeedGetter(feed *Feed) *feedGetter {
	getter := new(feedGetter)
	getter.feed = feed
	getter.feedConnector = rss.NewWithHandlers(5, true, nil, getter)
	return getter
}

func (getter *feedGetter) retrieveNewItems() ([]*rss.Item, error) {
	if err := getter.feedConnector.Fetch(getter.feed.Url, nil); err != nil {
		return nil, err
	}
	newItems := make([]*rss.Item, 0, len(getter.items))
	for _, rssItem := range getter.items {
		if pubDate, _ := rssItem.ParsedPubDate(); pubDate.After(getter.feed.lastSync) {
			newItems = append(newItems, rssItem)
		}
	}
	return newItems, nil
}

func (getter *feedGetter) ProcessItems(f *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	getter.items = append(getter.items, newitems...)
}
