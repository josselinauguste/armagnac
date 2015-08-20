package query

import "github.com/josselinauguste/armagnac/feeds/domain"
import "github.com/josselinauguste/armagnac/feeds/repository"

type NewItemsQuery struct {
	NewItems map[domain.Feed][]domain.Item
}

type message struct {
	feed  domain.Feed
	items []domain.Item
}

func NewNewItemsQuery() *NewItemsQuery {
	return &NewItemsQuery{}
}

func (query *NewItemsQuery) Execute() error {
	feeds := repository.CurrentFeedRepository.GetAll()
	query.updateQueryFeedsItems(feeds)
	return nil
}

func (query *NewItemsQuery) updateQueryFeedsItems(feeds []*domain.Feed) {
	ch := make(chan message)
	for _, feed := range feeds {
		go query.getFeedItems(feed, ch)
	}
	query.NewItems = make(map[domain.Feed][]domain.Item)
	for i := 0; i < len(feeds); i++ {
		r := <-ch
		if r.items != nil {
			query.NewItems[r.feed] = r.items
		}
	}
}

func (query *NewItemsQuery) getFeedItems(feed *domain.Feed, ch chan message) {
	getter := newFeedGetter(feed)
	newItems, err := getter.getNewItems()
	if err != nil {
		//TODO log
		ch <- message{}
	} else {
		ch <- message{*feed, newItems}
	}
}
