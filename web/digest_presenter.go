package web

import "github.com/josselinauguste/armagnac/feeds/query"

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
