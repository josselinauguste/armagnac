package repository

import "github.com/josselinauguste/armagnac/feeds/domain"

type feedRepository interface {
	GetAll() []*domain.Feed
	Persist(feed *domain.Feed)
}

var CurrentFeedRepository feedRepository
