package repository

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type feedDao struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Uri      string
	LastSync time.Time
}
