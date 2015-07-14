package feeds

import rss "github.com/jteeuwen/go-pkg-rss"

type FeedGetter struct {
	feed          *Feed
	feedConnector *rss.Feed
	items         []*rss.Item
}

func NewFeedGetter(feed *Feed) *FeedGetter {
	retriever := new(FeedGetter)
	retriever.feed = feed
	retriever.feedConnector = rss.NewWithHandlers(5, true, nil, retriever)
	return retriever
}

func (retriever *FeedGetter) RetrieveNewItems() ([]*rss.Item, error) {
	if err := retriever.feedConnector.Fetch(retriever.feed.url, nil); err != nil {
		return nil, err
	}
	newItems := make([]*rss.Item, 0, len(retriever.items))
	for _, item := range retriever.items {
		if pubDate, _ := item.ParsedPubDate(); pubDate.After(retriever.feed.lastSync) {
			newItems = append(newItems, item)
		}
	}
	return newItems, nil
}

func (retriever *FeedGetter) ProcessItems(f *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	retriever.items = append(retriever.items, newitems...)
}
