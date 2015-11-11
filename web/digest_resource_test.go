package web

import (
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
}

func (m *FakeMailer) sendMail(recipient string, subject string, content []byte) error {
	args := m.Called(recipient, subject, nil)
	return args.Error(0)
}

func TestCreateAndSendDigest(t *testing.T) {
	fakeMailer := new(FakeMailer)
	fakeMailer.On("sendMail", "jauguste@iblop.net", "A week digested", nil).Return(nil)
	fakeBus := new(FakeBus)
	fakeBus.On("Send", query.NewNewItemsQuery()).Return(nil)
	resource := newDigestResource(fakeBus, fakeMailer)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusOK)
	fakeMailer.AssertExpectations(t)
}
