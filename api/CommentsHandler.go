package api

import (
	"encoding/json"
	"fmt"
	"log"
	"matthewhope/go-webapp-js-components/services"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

var (
	cwg sync.WaitGroup
)

type CommentsHandler struct {
}

func NewCommentsHandler() *CommentsHandler {
	return new(CommentsHandler)
}

func (h *CommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CommentsHandler) get(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	c := make(chan services.Item)
	parent, exists := services.TopStories[id]
	if !exists {
		parent = services.GetItem(id)
	}
	fmt.Println(len(parent.Kids))
	var ret []services.Item
	cwg.Add(1)
	ct := 0
	go func() {
		for next := range c {
			ret = append(ret, next)
			ct++
			fmt.Println(ct)
			if ct >= len(parent.Kids) {
				cwg.Done()
				return
			}
		}
	}()
	services.GetItems(parent.Kids, c)
	cwg.Wait()
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Time > ret[j].Time
	})
	// fmt.Println(ret)
	err = json.NewEncoder(w).Encode(&ret)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
