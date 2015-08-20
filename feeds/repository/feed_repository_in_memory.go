// +build !appengine

package repository

import "github.com/josselinauguste/armagnac/feeds/domain"

type feedRepositoryInMemory struct {
	feeds []*domain.Feed
}

func newFeedRepositoryInMemory() feedRepository {
	return &feedRepositoryInMemory{make([]*domain.Feed, 0)}
}

func (repository *feedRepositoryInMemory) GetAll() []*domain.Feed {
	return repository.feeds
}

func (repository *feedRepositoryInMemory) Persist(feed *domain.Feed) {
	for _, storedFeed := range repository.feeds {
		if feed == storedFeed {
			return
		}
	}
	repository.feeds = append(repository.feeds, feed)
}

func init() {
	CurrentFeedRepository = newFeedRepositoryInMemory()
}
