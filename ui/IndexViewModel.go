package ui

import "matthewhope/go-webapp-js-components/services"

type IndexViewModel struct {
	Title string
	Items []services.Item
}

func NewIndexViewModel(title string) *IndexViewModel {
	ret := new(IndexViewModel)
	ret.Title = title
	ret.Items = services.GetTopStories()
	return ret
}
