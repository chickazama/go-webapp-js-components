package main

import (
	"log"
	"matthewhope/go-webapp-js-components/ui"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	mux.Handle("/", &ui.IndexHandler{})
	err := http.ListenAndServe(":8192", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
