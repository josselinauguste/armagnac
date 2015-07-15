// +build !appengine

package feeds

type feedRepositoryInMemory struct {
	feeds []*Feed
}

func newFeedRepositoryInMemory() feedRepository {
	return &feedRepositoryInMemory{make([]*Feed, 0)}
}

func (repository *feedRepositoryInMemory) GetAll() []*Feed {
	return repository.feeds
}

func (repository *feedRepositoryInMemory) Persist(feed *Feed) {
	for _, storedFeed := range repository.feeds {
		if feed == storedFeed {
			return
		}
	}
	repository.feeds = append(repository.feeds, feed)
}

func init() {
	currentFeedRepository = newFeedRepositoryInMemory()
}
