package domain

import "time"

type Feed struct {
	Uri      string
	LastSync time.Time
}

func NewFeed(uri string) *Feed {
	return &Feed{uri, time.Time{}}
}
