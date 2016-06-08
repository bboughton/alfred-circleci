package commands

import (
	"github.com/bboughton/alfred-circleci/circle"
	"github.com/bboughton/alfred-circleci/cli"
)

type Loadcache struct {
	Circle *circle.Client
}

func (c Loadcache) Exec(out cli.OutputWriter, in *cli.Input) {
	c.Circle.FindProjects("")
}
