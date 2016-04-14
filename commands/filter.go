package commands

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/bboughton/alfred-circleci/alfred"
	"github.com/bboughton/alfred-circleci/circle"
	"github.com/bboughton/alfred-circleci/cli"
	"github.com/bboughton/alfred-circleci/filter"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

type Filter struct {
	Circle *circle.Client
}

func (f Filter) Exec(out cli.OutputWriter, in *cli.Input) {
	args := strings.Split(in.Args[0], " ")

	var name string
	if len(args) > 0 {
		name = args[0]
	}

	resp := alfred.NewResponse()
	if !authenticated() {
		var token string
		if len(args) > 1 {
			token = args[1]
		}
		resp.AddItem(alfred.Item{
			Title: "/login",
			Arg:   "login " + token,
		})
	} else if string([]rune(name)[0]) == "/" {
		cmds := QueryCommands{
			{
				Title: "/logout",
				Arg:   "logout",
			},
			{
				Title: "/clear-cache",
				Arg:   "clearcache",
			},
		}

		filter.Filter(name, &cmds, fuzzy.Match)

		for _, cmd := range cmds {
			resp.AddItem(alfred.Item{
				Title: cmd.Title,
				Arg:   cmd.Arg,
			})
		}
	} else if len(name) > 0 {
		projects := f.Circle.FindProjects(name)
		for _, proj := range projects {
			resp.AddItem(alfred.Item{
				Title: proj.Name,
				Arg:   "open " + proj.URL,
			})
		}
	}

	fmt.Print(xml.Header)
	if err := xml.NewEncoder(os.Stdout).Encode(resp); err != nil {
		fmt.Println(err)
		out.ExitWith(1)
	}
}

func authenticated() bool {
	file, err := os.Open("auth.json")
	file.Close()
	return !os.IsNotExist(err)
}

type QueryCommand struct {
	Title string
	Arg   string
}

type QueryCommands []QueryCommand

func (c QueryCommands) Len() int {
	return len(c)
}

func (c QueryCommands) Index(i int) string {
	return c[i].Title
}

func (c *QueryCommands) Remove(i int) {
	*c = append([]QueryCommand(*c)[:i], []QueryCommand(*c)[i+1:]...)
}

func (c *QueryCommands) Add(cmd QueryCommand) {
	*c = append(*c, cmd)
}
