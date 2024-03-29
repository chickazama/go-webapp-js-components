package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJobStoriesIDs() ([]int64, error) {
	return getStoriesIDs(jobStoriesEndpoint)
}

func GetNewStoriesIDs() ([]int64, error) {
	return getStoriesIDs(newStoriesEndpoint)
}

func GetTopStoriesIDs() ([]int64, error) {
	return getStoriesIDs(topStoriesEndpoint)
}

func getStoriesIDs(endpoint string) ([]int64, error) {
	var ret []int64
	url := fmt.Sprintf("%s%s", baseUrl, endpoint)
	res, err := fetch(http.MethodGet, url, nil)
	if err != nil {
		return ret, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
