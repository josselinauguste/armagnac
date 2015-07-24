package web

import (
	"net/http"
	"strings"
	"testing"

	"github.com/josselinauguste/armagnac/feeds"
	"github.com/josselinauguste/magicbus"

	"github.com/stretchr/testify/mock"
)

type FakeBus struct {
	mock.Mock
}

func (m *FakeBus) Send(command magicbus.Command) error {
	m.Called(command)
	return nil
}

func TestCreateAndSendDigest(t *testing.T) {
	fakeBus := new(FakeBus)
	fakeBus.On("Send", feeds.NewNewItemsQuery()).Return(nil)
	resource := newDigestResource(fakeBus)
	request, _ := http.NewRequest("POST", "/digests", strings.NewReader(``))
	response := NewFakeResponse(t)

	resource.createAndSendDigestHandler(response, request)

	response.AssertStatus(http.StatusOK)
	fakeBus.AssertExpectations(t)
}
