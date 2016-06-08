package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/bboughton/alfred-circleci/circle"
	"github.com/bboughton/alfred-circleci/cli"
	"github.com/bboughton/alfred-circleci/commands"
)

const (
	FOUR_HOURS = 14400
)

func main() {
	var (
		// path to cache file
		cachePath string

		// path to auth file
		authPath string

		// how long to keep cache
		ttl        time.Duration
		ttlSeconds int
	)

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	flag.StringVar(&cachePath, "cache", fmt.Sprintf("%s/.cache/alfred-circleci/cache.json", usr.HomeDir), "path to cache file")
	flag.StringVar(&authPath, "auth", fmt.Sprintf("%s/.config/alfred-circleci/auth.json", usr.HomeDir), "path to auth file")
	flag.IntVar(&ttlSeconds, "ttl", FOUR_HOURS, "number of seconds to keep the cache")
	flag.Parse()

	ttl = time.Duration(ttlSeconds) * time.Second

	cmd := cli.NewSubCommandHandler()

	// Filter command
	auth := LoadAuth(authPath)
	client := circle.NewClient(auth.Token, cachePath, ttl)
	cmd.Handle("filter", AuthHandler{
		Auth: auth,
		Handler: commands.Filter{
			Circle: client,
		},
	})

	// Run Command
	run := cli.NewSubCommandHandler()
	run.Handle("open", commands.Open{})
	run.Handle("login", commands.Login{
		AuthPath: authPath,
	})
	run.Handle("loadcache", AuthHandler{
		Auth: auth,
		Handler: commands.Loadcache{
			Circle: client,
		},
	})
	run.Handle("clearcache", commands.Clearcache{
		CachePath: cachePath,
	})
	run.Handle("logout", commands.Logout{
		AuthPath: authPath,
	})
	cmd.Handle("run", run)

	app := cli.New(cmd)
	os.Exit(app.Run(flag.Args()))
}
