package feeds

import "time"

type Feed struct {
	Uri      string
	lastSync time.Time
}

func NewFeed(uri string) *Feed {
	return &Feed{uri, time.Time{}}
}
