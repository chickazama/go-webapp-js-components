package services

import (
	"io"
	"net/http"
)

const (
	baseUrl            = "https://hacker-news.firebaseio.com/v0/"
	maxItemEndpoint    = "maxitem.json"
	newStoriesEndpoint = "newstories.json"
)

func fetch(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}
