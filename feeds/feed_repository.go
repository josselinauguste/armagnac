package feeds

type feedRepository interface {
	GetAll() []*Feed
	Add(feed *Feed)
}

var currentFeedRepository feedRepository
