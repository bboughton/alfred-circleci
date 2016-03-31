package main

import (
	"fmt"
	"os"
)

type InvalidCommand struct {
}

func (c InvalidCommand) Run(args []string) int {
	fmt.Fprintln(os.Stderr, "invalid action:", args)
	return 1
}
