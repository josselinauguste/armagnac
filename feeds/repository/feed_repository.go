package repository

import "github.com/josselinauguste/armagnac/feeds/domain"

type feedRepository interface {
	GetAll() ([]*domain.Feed, error)
	Persist(feed *domain.Feed) error
}

var CurrentFeedRepository feedRepository
