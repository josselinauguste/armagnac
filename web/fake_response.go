package web

import (
	"encoding/json"
	"net/http"
	"testing"
)

type FakeResponse struct {
	t       *testing.T
	headers http.Header
	body    []byte
	status  int
}

func NewFakeResponse(t *testing.T) *FakeResponse {
	return &FakeResponse{
		t:       t,
		headers: make(http.Header),
	}
}

func (r FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	if r.status == 0 {
		r.WriteHeader(http.StatusOK)
	}
	r.body = body
	if r.headers["Content-Type"] == nil {
		r.Header().Set("Content-Type", http.DetectContentType(body))
	}
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func (r FakeResponse) GetBody() string {
	return string(r.body)
}

func (r FakeResponse) GetJsonBody(v interface{}) error {
	return json.Unmarshal(r.body, v)
}

func (r FakeResponse) AssertStatus(status int) {
	if r.status != status {
		r.t.Errorf("expected status %+v to equal %+v", r.status, status)
	}
}
