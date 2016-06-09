package alfred

import (
	"encoding/json"
	"io"
)

type Response struct {
	Items []Item `json:"items"`
}

func NewResponse() *Response {
	return &Response{
		Items: make([]Item, 0),
	}
}

func (r *Response) AddItem(item Item) {
	r.Items = append(r.Items, item)
}

type Item struct {
	UID          string `json:"uid"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Icon         Icon   `json:"icon"`
	Valid        bool   `json:"valid"`
	Autocomplete string `json:"autocomplete"`
	Type         Type   `json:"type"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Type string

const (
	DefaultType       Type = "default"
	FileType               = "file"
	FileSkipCheckType      = "file:skipcheck"
)

func NewItem(title, arg string) Item {
	return Item{
		UID:          title,
		Title:        title,
		Arg:          arg,
		Valid:        true,
		Autocomplete: title,
		Type:         DefaultType,
	}
}

func WriteResponse(w io.Writer, resp *Response) error {
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}
	return nil
}
