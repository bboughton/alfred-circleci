package commands

import (
	"os"

	"github.com/bboughton/alfred-circleci/cli"
)

type Clearcache struct {
	CachePath string
}

func (c Clearcache) Exec(out cli.OutputWriter, in *cli.Input) {
	os.Remove(c.CachePath)
}
