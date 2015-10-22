// +build mongodb

package repository

import (
	"fmt"
	"os"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/josselinauguste/armagnac/feeds/domain"
)

type feedRepositoryMongoDb struct {
	initialSession *mgo.Session
}

func newFeedRepositoryMongoDb() feedRepository {
	return &feedRepositoryMongoDb{}
}

func (repository *feedRepositoryMongoDb) acquireSession() (*mgo.Session, error) {
	session, err := mgo.Dial(os.Getenv("MONGODB_URI")) //TODO share session by creating other ones with session.Copy()
	if err != nil {
		fmt.Println("ERROR: can't open MongoDB connection: %#v", err.Error())
		return nil, err
	}
	session.SetSafe(&mgo.Safe{})
	repository.initialSession = session
	return session, nil
}

func (repository *feedRepositoryMongoDb) GetAll() ([]*domain.Feed, error) {
	session, err := repository.acquireSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	collection := session.DB("armagnac").C("feeds")
	var feeds []*domain.Feed
	var feedDao feedDao
	iter := collection.Find(nil).Iter()
	for iter.Next(&feedDao) {
		feed := mapToFeed(feedDao)
		feeds = append(feeds, &feed)
	}
	if err := iter.Close(); err != nil {
		fmt.Println("ERROR: can't close MongoDB iterator: %#v", err.Error())
		return feeds, err
	}
	return feeds, nil
}

func mapToFeed(feed feedDao) domain.Feed {
	return domain.Feed{feed.ID.Hex(), "", feed.Uri, feed.LastSync}
}

func (repository *feedRepositoryMongoDb) Persist(feed *domain.Feed) error {
	session, err := repository.acquireSession()
	if err != nil {
		return err
	}
	defer session.Close()
	collection := session.DB("armagnac").C("feeds")
	feedDao := mapFromFeed(feed)
	_, err = collection.UpsertId(feedDao.ID, bson.M{"$set": feedDao})
	return err
}

func mapFromFeed(feed *domain.Feed) feedDao {
	var id bson.ObjectId
	if len(feed.ID) > 0 {
		id = bson.ObjectIdHex(feed.ID)
	} else {
		id = bson.NewObjectId()
		feed.ID = id.Hex()
	}
	return feedDao{id, feed.Uri, feed.LastSync}
}

func (repository *feedRepositoryMongoDb) RemoveAll() error {
	return nil
}

func init() {
	fmt.Println("INFO: Using MongoDB repository")
	CurrentFeedRepository = newFeedRepositoryMongoDb()
}
