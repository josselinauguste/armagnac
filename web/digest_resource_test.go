package web

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
	"testing"

	"github.com/josselinauguste/armagnac/feeds/query"
	"github.com/josselinauguste/magicbus"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/josselinauguste/armagnac/feeds/domain"
	"time"
)

type FakeBus struct {
	mock.Mock
	mockedFeeds []*domain.Feed
	mockedNewItems map[string][]domain.Item
}

func (m *FakeBus) Send(command magicbus.Command) error {
	args := m.Called(command)
	query := command.(*query.NewItemsQuery)
	query.Feeds = m.mockedFeeds
	query.NewItems = m.mockedNewItems
	return args.Error(0)
}

type FakeMailer struct {
	mock.Mock
	content []byte
}

func (m *FakeMailer) sendMail(recipient string, subject string, content []byte) error {
	m.content = content
	args := m.Called(recipient, subject, content)
	return args.Error(0)
}

func TestCreateAndSendDigest(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeMailer.On("sendMail", "jauguste@iblop.net", "A week digested", mock.AnythingOfType("[]uint8")).Return(nil)
	feeds, items := createData()
	fakeBus := new(FakeBus)
	fakeBus.mockedFeeds = feeds
	fakeBus.mockedNewItems = items
	fakeBus.On("Send", mock.AnythingOfType("*query.NewItemsQuery")).Return(nil)
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusOK)
	fakeBus.AssertExpectations(t)
	fakeMailer.AssertExpectations(t)
	digest := string(fakeMailer.content)
	item := items[feeds[0].ID][0]
	assert.Contains(t, digest, item.Title)
	assert.Contains(t, digest, item.Description)
}

func createData() ([]*domain.Feed, map[string][]domain.Item) {
	items := make(map[string][]domain.Item)
	feed := domain.NewFeed("http://test.com")
	items[feed.ID] = make([]domain.Item, 0)
	item := domain.Item{
		Title:           "title",
		Url:             "http://url.com",
		Description:     "description",
		PublicationDate: time.Now(),
	}
	items[feed.ID] = append(items["id"], item)
	feeds := make([]*domain.Feed, 0)
	feeds = append(feeds, feed)
	return  feeds, items
}

func TestBusErrorReturns500(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeBus := new(FakeBus)
	fakeBus.On("Send", mock.AnythingOfType("*query.NewItemsQuery")).Return(errors.New(""))
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusInternalServerError)
}

func TestErrorOnEmailSendReturns500(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeMailer.On("sendMail", "jauguste@iblop.net", "A week digested", mock.AnythingOfType("[]uint8")).Return(errors.New(""))
	fakeBus := new(FakeBus)
	fakeBus.On("Send", mock.AnythingOfType("*query.NewItemsQuery")).Return(nil)
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusInternalServerError)
}

func TestErrorOnTemplatingReturns500(t *testing.T) {
	digestTemplate, _ = template.New("test").Parse("{{.Count}} is not a valid template")
	fakeMailer := new(FakeMailer)
	fakeMailer.On("sendMail", "jauguste@iblop.net", "A week digested", nil).Return(nil)
	fakeBus := new(FakeBus)
	fakeBus.On("Send", mock.AnythingOfType("*query.NewItemsQuery")).Return(nil)
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusInternalServerError)
}
