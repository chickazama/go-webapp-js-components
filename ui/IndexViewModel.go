package ui

type IndexViewModel struct {
	Title string
}

func NewIndexViewModel(title string) *IndexViewModel {
	ret := new(IndexViewModel)
	ret.Title = title
	return ret
}
