package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func GetNewItems(old, new int64, c chan Item) {
	fmt.Printf("MAX ID: %d\n", new)
	for i := old + 1; i <= new; i++ {
		id := i
		wg.Add(1)
		go func() {
			GetItemAsync(id, c)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Batch done.")
}

func GetItems(ids []int64, c chan Item) {
	for _, i := range ids {
		id := i
		wg.Add(1)
		go func() {
			GetItemAsync(id, c)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Batch done.")
}

func GetItemAsync(id int64, c chan Item) {
	url := fmt.Sprintf("%sitem/%d.json", baseUrl, id)
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
			// if out.Type == "story" {
			c <- out
			// }
			break
		} else {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Error")
		}
	}
}

func GetItem(id int64) Item {
	url := fmt.Sprintf("%sitem/%d.json", baseUrl, id)
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
			return out
		} else {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Error")
		}
	}
}
