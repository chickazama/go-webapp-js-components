package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetItems(maxItemID, old int64) {
	for i := maxItemID; i > old; i-- {
		url := fmt.Sprintf("%s/item/%d.json", baseUrl, i)
		res, err := fetch(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		var out Item
		err = json.Unmarshal(buf, &out)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("%+v\n", out)
	}
}
