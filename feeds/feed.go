package feeds

type Feed struct {
  url string
}

func NewFeed(url string) *Feed {
  return &Feed{url}
}
