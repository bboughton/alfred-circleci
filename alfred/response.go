package alfred

import "encoding/xml"

type Response struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) AddItem(item Item) {
	r.Items = append(r.Items, item)
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Arg     string   `xml:"arg"`
}

func NewItem(title string) Item {
	return Item{
		Title: title,
	}
}
