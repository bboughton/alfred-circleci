package commands

import (
	"fmt"

	"github.com/bboughton/alfred-circleci/circle"
	"github.com/bboughton/alfred-circleci/cli"
)

type Loadcache struct {
	Circle *circle.Client
}

func (c Loadcache) Exec(out cli.OutputWriter, in *cli.Input) {
	_, err := c.Circle.FindProjects("")
	if err != nil {
		fmt.Fprintln(out, err)
		out.ExitWith(1)
	}
}
