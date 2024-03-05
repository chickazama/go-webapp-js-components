package services

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetMaxItemID() (int64, error) {
	url := fmt.Sprintf("%s%s", baseUrl, maxItemEndpoint)
	res, err := fetch(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
