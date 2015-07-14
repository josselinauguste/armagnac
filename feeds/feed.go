package feeds

import "time"

type Feed struct {
	url      string
	lastSync time.Time
}

func NewFeed(url string) *Feed {
	return &Feed{url, time.Time{}}
}
