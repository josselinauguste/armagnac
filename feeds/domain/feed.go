package domain

import "time"

type Feed struct {
	ID       string
	Title    string
	Uri      string
	LastSync time.Time
}

func NewFeed(uri string) *Feed {
	return &Feed{"", "", uri, time.Time{}}
}
