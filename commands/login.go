package commands

import (
	"fmt"
	"os"

	"github.com/bboughton/alfred-circleci/cli"
)

type Login struct {
	AuthPath string
}

func (l Login) Exec(out cli.OutputWriter, in *cli.Input) {
	if len(in.Args) == 0 {
		fmt.Fprintln(out, "token required")
		out.ExitWith(1)
		return
	}

	token := in.Args[0]
	err := SaveAuth(l.AuthPath, Auth{Token: token})
	if err != nil {
		os.Remove(l.AuthPath)
		fmt.Fprintln(out, err)
		out.ExitWith(1)
	}
}
