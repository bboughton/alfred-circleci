package circle

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bboughton/alfred-circleci/circle/api"
	"github.com/bboughton/alfred-circleci/circle/cache"
)

// The RepoLoader interface should be used to load repos
type ProjectLoader interface {
	LoadProject() (Projects, error)
}

// CacheRepoLoader is a RepoLoader that will cache its source to the given path
type CacheProjectLoader struct {
	Path   string
	TTL    time.Duration
	Source ProjectLoader
}

// LoadRepos will load repos from the cache unless it has expired in which case
// it will retreive the repos from the Source and update the cache.
func (loader CacheProjectLoader) LoadProject() (Projects, error) {
	var (
		projs Projects
		err   error
	)

	item := cache.Load(loader.Path)
	if item.Expired() {
		projs, err = loader.Source.LoadProject()
		if err != nil {
			return projs, err
		}

		data, err := json.Marshal(projs)
		if err != nil {
			return projs, err
		}
		cache.Save(loader.Path, cache.NewItem(loader.TTL, data))
	} else {
		json.Unmarshal(item.Data, &projs)
	}

	return projs, nil
}

// APIRepoLoader is a RepoLoader that will retrive the repos from the API
type APIProjectLoader struct {
	Client *api.Client
}

// LoadRepos will load all repos from the api by itterating over each page
// of repos until there are no more to consume
func (loader APIProjectLoader) LoadProject() (Projects, error) {
	var (
		projects Projects
		repos    []api.Repo
		page     int = 1
		err      error
	)

	for hasNextPage := true; hasNextPage; hasNextPage = (len(repos) > 0) {
		repos, err = loader.Client.ListRepos(page)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			projects.Add(Project{
				Name: fmt.Sprintf("%s/%s", repo.Username, repo.Name),
				URL:  fmt.Sprintf("https://circleci.com/gh/%s/%s", repo.Username, repo.Name),
			})
		}
		page++
	}

	return projects, nil
}
