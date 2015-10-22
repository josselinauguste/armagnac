// +build !mongodb

package repository

import (
	"fmt"

	"github.com/josselinauguste/armagnac/feeds/domain"
)

type feedRepositoryInMemory struct {
	feeds []*domain.Feed
}

func newFeedRepositoryInMemory() feedRepository {
	return &feedRepositoryInMemory{make([]*domain.Feed, 0)}
}

func (repository *feedRepositoryInMemory) GetAll() ([]*domain.Feed, error) {
	return repository.feeds, nil
}

func (repository *feedRepositoryInMemory) Persist(feed *domain.Feed) error {
	for _, storedFeed := range repository.feeds {
		if feed == storedFeed {
			return nil
		}
	}
	repository.feeds = append(repository.feeds, feed)
	return nil
}

func (repository *feedRepositoryInMemory) RemoveAll() error {
	repository.feeds = make([]*domain.Feed, 0)
	return nil
}

func init() {
	fmt.Println("INFO: Using in-memory repository")
	CurrentFeedRepository = newFeedRepositoryInMemory()
}
