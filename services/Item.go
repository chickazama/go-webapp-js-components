package services

type Item struct {
	ID          int64   `json:"id,omitempty"`
	Author      string  `json:"by,omitempty"`
	Dead        bool    `json:"dead,omitempty"`
	Descendants int64   `json:"descendants,omitempty"`
	Deleted     bool    `json:"deleted,omitempty"`
	Kids        []int64 `json:"kids,omitempty"`
	Parent      int64   `json:"parent,omitempty"`
	Parts       []int64 `json:"parts,omitempty"`
	Poll        int64   `json:"poll,omitempty"`
	Score       int64   `json:"score,omitempty"`
	Text        string  `json:"text,omitempty"`
	Title       string  `json:"title,omitempty"`
	Time        int64   `json:"time,omitempty"`
	Type        string  `json:"type,omitempty"`
	URL         string  `json:"url,omitempty"`
}

func (i *Item) Comments() int {
	return len(i.Kids)
}
