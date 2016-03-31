package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c Client) ListRepos(page int) []Repo {
	var (
		host         string = "https://circleci.com"
		basePath     string = "/api/v1"
		resourcePath string = "/user/repos"
		err          error
	)

	req, err := http.NewRequest(http.MethodGet, host+basePath+resourcePath, nil)
	if err != nil {
		return nil
	}

	values := make(url.Values)
	values.Set("circle-token", c.Token)
	values.Set("page", fmt.Sprintf("%d", page))
	req.URL.RawQuery = values.Encode()

	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var repos []Repo
	if err = json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil
	}

	return repos
}
