package feeds

type NewItemsQuery struct {
	NewItems map[Feed][]Item
}

type message struct {
	feed  Feed
	items []Item
}

func NewNewItemsQuery() *NewItemsQuery {
	return &NewItemsQuery{}
}

func (query *NewItemsQuery) Execute() error {
	feeds := currentFeedRepository.GetAll()
	query.updateQueryFeedsItems(feeds)
	return nil
}

func (query *NewItemsQuery) updateQueryFeedsItems(feeds []*Feed) {
	ch := make(chan message)
	for _, feed := range feeds {
		go query.getFeedItems(feed, ch)
	}
	query.NewItems = make(map[Feed][]Item)
	for i := 0; i < len(feeds); i++ {
		r := <-ch
		if r.items != nil {
			query.NewItems[r.feed] = r.items
		}
	}
}

func (query *NewItemsQuery) getFeedItems(feed *Feed, ch chan message) {
	getter := newFeedGetter(feed)
	newItems, err := getter.getNewItems()
	if err != nil {
		//TODO log
		ch <- message{}
	} else {
		ch <- message{*feed, newItems}
	}
}
