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

func (repository *feedRepositoryInMemory) Add(feed *Feed) {
	repository.feeds = append(repository.feeds, feed)
}

func init() {
	currentFeedRepository = newFeedRepositoryInMemory()
}
