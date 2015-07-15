package feeds

type feedRepository interface {
	GetAll() []*Feed
	Persist(feed *Feed)
}

var currentFeedRepository feedRepository
