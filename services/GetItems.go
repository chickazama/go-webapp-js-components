package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func GetItems(old, new int64, c chan int64) {
	fmt.Printf("MAX ID: %d\n", new)
	for i := old + 1; i <= new; i++ {
		url := fmt.Sprintf("%s/item/%d.json", baseUrl, i)
		wg.Add(1)
		go func() {
			for {
				res, err := fetch(http.MethodGet, url, nil)
				if err != nil {
					log.Fatal(err.Error())
				}
				buf, err := io.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err.Error())
				}
				res.Body.Close()
				var out Item
				err = json.Unmarshal(buf, &out)
				if err != nil {
					log.Fatal(err.Error())
				}
				if out.ID != 0 {
					fmt.Printf("[%d]: %s\n", out.ID, out.Type)
					if out.Type == "story" {
						c <- out.ID
					}
					break
				} else {
					fmt.Println("Miss. Retrying...")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
