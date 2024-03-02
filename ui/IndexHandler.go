package ui

import (
	"html/template"
	"net/http"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return new(IndexHandler)
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *IndexHandler) get(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./templates/*.go.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vm := NewIndexViewModel("Home Page")
	err = tmpl.ExecuteTemplate(w, "Index", vm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
