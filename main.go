package main

import (
	"os"
	"time"

	"github.com/bboughton/alfred-circleci/circle"
	"github.com/bboughton/alfred-circleci/cli"
	"github.com/bboughton/alfred-circleci/commands"
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
	cmd := cli.NewSubCommandHandler()

	// Filter command
	auth := LoadAuth(authPath)
	client := circle.NewClient(auth.Token, cachePath, ttl)
	cmd.Handle("filter", commands.Filter{
		Circle: client,
	})

	// Run Command
	run := cli.NewSubCommandHandler()
	run.Handle("open", commands.Open{})
	run.Handle("login", commands.Login{
		AuthPath: authPath,
	})
	run.Handle("loadcache", commands.Loadcache{
		Circle: client,
	})
	run.Handle("clearcache", commands.Clearcache{
		CachePath: cachePath,
	})
	run.Handle("logout", commands.Logout{
		AuthPath: authPath,
	})
	cmd.Handle("run", run)

	os.Exit(cli.Run(cmd))
}
