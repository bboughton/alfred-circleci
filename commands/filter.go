package commands

import (
	"fmt"
	"os"

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
	var name string
	if len(in.Args) > 0 {
		name = in.Args[0]
	}

	resp := alfred.NewResponse()
	if len(name) > 0 && string([]rune(name)[0]) == "/" {
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
			resp.AddItem(alfred.NewItem(cmd.Title, cmd.Arg))
		}
	} else if len(name) > 0 {
		projects := f.Circle.FindProjects(name)
		for _, proj := range projects {
			resp.AddItem(alfred.NewItem(proj.Name, "open "+proj.URL))
		}
	}

	err := alfred.WriteResponse(os.Stdout, resp)
	if err != nil {
		fmt.Println(err)
		out.ExitWith(1)
	}
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
