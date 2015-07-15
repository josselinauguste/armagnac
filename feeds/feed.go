package feeds

import "time"

type Feed struct {
	Uri      string
	lastSync time.Time
}

func newFeed(uri string) *Feed {
	return &Feed{uri, time.Time{}}
}
