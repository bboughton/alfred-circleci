package circle

import (
	"time"

	"github.com/bboughton/alfred-circleci/circle/api"
	"github.com/bboughton/alfred-circleci/filter"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// Client is used to access the circle api
type Client struct {
	loader ProjectLoader
}

// New creates a new Client
func NewClient(token, path string, ttl time.Duration) *Client {
	return &Client{
		loader: CacheProjectLoader{
			Path: path,
			TTL:  ttl,
			Source: APIProjectLoader{
				Client: api.NewClient(token),
			},
		},
	}
}

func (c Client) FindProjects(name string) (Projects, error) {
	// load all repos and map them to projects
	all, err := c.loader.LoadProject()
	if err != nil {
		return nil, err
	}

	filter.Filter(name, &all, fuzzy.Match)

	return all, nil
}
