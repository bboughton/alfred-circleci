package commands

import (
	"os"

	"github.com/bboughton/alfred-circleci/cli"
)

type Logout struct {
	AuthPath string
}

func (l Logout) Exec(out cli.OutputWriter, in *cli.Input) {
	os.Remove(l.AuthPath)
}
