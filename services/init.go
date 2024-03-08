package services

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

var (
	jobStories       map[int64]Item
	newStories       map[int64]Item
	TopStories       map[int64]Item
	sortedJobStories []Item
	sortedNewStories []Item
	sortedTopStories []Item
)

var (
	jsMtx  sync.Mutex
	nsMtx  sync.Mutex
	tsMtx  sync.Mutex
	initWg sync.WaitGroup
)

func Init() {
	jobStories = make(map[int64]Item)
	newStories = make(map[int64]Item)
	TopStories = make(map[int64]Item)
	initWg.Add(3)
	go func() {
		defer initWg.Done()
		err := populate(TopStories, GetTopStoriesIDs)
		if err != nil {
			log.Fatal(err.Error())
		}
		sortedTopStories = setTopStories()
	}()
	go func() {
		defer initWg.Done()
		err := populate(newStories, GetNewStoriesIDs)
		if err != nil {
			log.Fatal(err.Error())
		}
		sortedNewStories = setNewStories()
	}()
	go func() {
		defer initWg.Done()
		err := populate(jobStories, GetJobStoriesIDs)
		if err != nil {
			log.Fatal(err.Error())
		}
		sortedJobStories = setJobStories()
	}()
	initWg.Wait()
	fmt.Println("Done.")
}

func populate(m map[int64]Item, f func() ([]int64, error)) error {
	var wg sync.WaitGroup
	ids, err := f()
	if err != nil {
		return err
	}
	// fmt.Println(len(ids))
	c := make(chan Item)
	ct := 0
	wg.Add(1)
	go func() {
		for item := range c {
			m[item.ID] = item
			ct++
			// fmt.Println(ct)
			// fmt.Printf("Completed: %d\n", ct)
			if ct >= len(ids) {
				wg.Done()
				return
			}
		}
	}()
	GetItems(ids, c)
	wg.Wait()
	return nil
}

func setTopStories() []Item {
	var ret []Item
	for _, v := range TopStories {
		ret = append(ret, v)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Score > ret[j].Score
	})
	return ret
}

func setNewStories() []Item {
	var ret []Item
	for _, v := range newStories {
		ret = append(ret, v)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Time > ret[j].Time
	})
	return ret
}

func setJobStories() []Item {
	var ret []Item
	for _, v := range jobStories {
		ret = append(ret, v)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Time > ret[j].Time
	})
	return ret
}

func fixedUpdate() {
	maxItemID, err := GetMaxItemID()
	if err != nil {
		log.Fatal(err.Error())
	}
	c := make(chan Item)
	GetNewItems(0, maxItemID, c)
}

func GetTopStories() []Item {
	tsMtx.Lock()
	ret := sortedTopStories
	tsMtx.Unlock()
	return ret
}

func GetJobStories() []Item {
	jsMtx.Lock()
	ret := sortedJobStories
	jsMtx.Unlock()
	return ret
}

func GetNewStories() []Item {
	nsMtx.Lock()
	ret := sortedNewStories
	nsMtx.Unlock()
	return ret
}
