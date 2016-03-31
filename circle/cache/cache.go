package cache

import (
	"encoding/json"
	"os"
	"time"
)

type Item struct {
	Data    []byte
	Expires time.Time
}

func NewItem(ttl time.Duration, data []byte) Item {
	return Item{
		Data:    data,
		Expires: time.Now().Add(ttl),
	}
}

func (c Item) Expired() bool {
	return time.Now().After(c.Expires)
}

type Data []byte

func Load(path string) Item {
	cache := Item{}

	file, err := os.Open(path)
	if err != nil {
		return cache
	}

	json.NewDecoder(file).Decode(&cache)
	return cache
}

func Save(path string, item Item) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if os.IsNotExist(err) {
		file, err = os.Create(path)
	}

	if err != nil {
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(item)
}
