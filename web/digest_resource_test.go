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
)

type FakeBus struct {
	mock.Mock
}

func (m *FakeBus) Send(command magicbus.Command) error {
	args := m.Called(command)
	return args.Error(0)
}

func TestCreateDigest(t *testing.T) {
	fakeBus := new(FakeBus)
	fakeBus.On("Send", query.NewNewItemsQuery()).Return(nil)
	resource := newDigestResource(fakeBus, new(FakeMailer))

	digest, err := resource.createDigest()

	assert.NotEmpty(t, digest)
	assert.Nil(t, err)
	fakeBus.AssertExpectations(t)
}

type FakeMailer struct {
	mock.Mock
	content []byte
}

func (m *FakeMailer) sendMail(recipient string, subject string, content []byte) error {
	m.content = content
	args := m.Called(recipient, subject, nil)
	return args.Error(0)
}

func TestCreateAndSendDigest(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeMailer.On("sendMail", "jauguste@iblop.net", "A week digested", nil).Return(nil)
	fakeBus := new(FakeBus)
	query := query.NewNewItemsQuery()
	fakeBus.On("Send", query).Return(nil)
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusOK)
	fakeBus.AssertExpectations(t)
	fakeMailer.AssertExpectations(t)
}

func TestBusErrorReturns500(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeBus := new(FakeBus)
	query := query.NewNewItemsQuery()
	fakeBus.On("Send", query).Return(errors.New(""))
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusInternalServerError)
}

func TestErrorOnEmailSendReturns500(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeMailer.On("sendMail", "jauguste@iblop.net", "A week digested", nil).Return(errors.New(""))
	fakeBus := new(FakeBus)
	query := query.NewNewItemsQuery()
	fakeBus.On("Send", query).Return(nil)
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
	query := query.NewNewItemsQuery()
	fakeBus.On("Send", query).Return(nil)
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusInternalServerError)
}
