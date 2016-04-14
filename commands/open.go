package commands

import (
	"fmt"
	"os/exec"

	"github.com/bboughton/alfred-circleci/cli"
)

type Open struct {
}

func (o Open) Exec(out cli.OutputWriter, in *cli.Input) {
	if len(in.Args) == 0 {
		fmt.Fprintln(out, "url required")
		out.ExitWith(1)
		return
	}

	url := in.Args[0]
	err := exec.Command("open", url).Run()
	if err != nil {
		fmt.Fprintln(out, err)
		out.ExitWith(1)
	}
}
