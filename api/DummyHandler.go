package api

import (
	"encoding/json"
	"net/http"
)

type DummyHandler struct {
}

type DummyPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewDummyHandler() *DummyHandler {
	return new(DummyHandler)
}

func (h *DummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p := &DummyPayload{Name: "Matt", Age: 27}
		err := json.NewEncoder(w).Encode(p)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
