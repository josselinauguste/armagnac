package feeds

import "time"

type Feed struct {
	Url      string
	lastSync time.Time
}

func newFeed(url string) *Feed {
	return &Feed{url, time.Time{}}
}
