package feeds

type NewItemsQuery struct {
	NewItems map[Feed][]Item
}

func (query *NewItemsQuery) Execute() {
	query.NewItems = make(map[Feed][]Item)
	feeds := currentFeedRepository.GetAll()
	for _, feed := range feeds {
		query.NewItems[*feed] = query.getFeedItems(feed)
	}
}

func (query *NewItemsQuery) getFeedItems(feed *Feed) []Item {
	getter := newFeedGetter(feed)
	newItems, err := getter.getNewItems()
	if err != nil {
		//TODO log
		return nil
	}
	return newItems
}
