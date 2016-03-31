package main

import (
	"os"
	"time"

	"github.com/bboughton/alfred-circleci/circle"
	"github.com/bboughton/alfred-circleci/cli"
)

const (
	// path to cache file
	cachePath string = "cache.json"

	// path to auth file
	authPath string = "auth.json"

	// how long to keep cache
	ttl time.Duration = 4 * time.Hour
)

func main() {
	cmd := cli.NewCommandHandler(InvalidCommand{})

	auth := LoadAuth(authPath)
	cmd.Handle("filter", FilterCommand{
		Circle: circle.NewClient(auth.Token, cachePath, ttl),
	})
	cmd.Handle("run", RunCommand{})

	os.Exit(cmd.Run(os.Args[1:]))
}
