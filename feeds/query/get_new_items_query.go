package query

import (
	"fmt"

	"github.com/josselinauguste/armagnac/feeds/domain"
)

import "github.com/josselinauguste/armagnac/feeds/repository"

type NewItemsQuery struct {
	Feeds    []*domain.Feed
	NewItems map[string][]domain.Item
}

type message struct {
	feed  *domain.Feed
	items []domain.Item
}

func NewNewItemsQuery() *NewItemsQuery {
	return &NewItemsQuery{nil, make(map[string][]domain.Item)}
}

func (query *NewItemsQuery) Execute() error {
	feeds, err := repository.CurrentFeedRepository.GetAll()
	if err != nil {
		return err
	}
	query.updateQueryFeedsItems(feeds)
	return nil
}

func (query *NewItemsQuery) updateQueryFeedsItems(feeds []*domain.Feed) {
	ch := make(chan message)
	for _, feed := range feeds {
		go query.getFeedItems(feed, ch)
	}
	for i := 0; i < len(feeds); i++ {
		r := <-ch
		if r.items != nil {
			query.Feeds = append(query.Feeds, r.feed)
			query.NewItems[r.feed.ID] = r.items
		}
	}
}

func (query *NewItemsQuery) getFeedItems(feed *domain.Feed, ch chan message) {
	getter := newFeedGetter(feed)
	newItems, err := getter.getNewItems()
	if err != nil {
		fmt.Println("ERROR: can't get feed %v: %#v", feed.Uri, err.Error())
		//TODO use mutex release instead of empty message
		ch <- message{}
	} else {
		ch <- message{feed, newItems}
	}
}
