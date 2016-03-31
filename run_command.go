package main

import (
	"fmt"
	"os"
	"os/exec"
)

type RunCommand struct {
}

func (c RunCommand) Run(args []string) int {
	if len(args) == 0 {
		return 1
	}
	action := args[0]

	switch action {
	case "open":
		err := exec.Command("open", args[1]).Run()
		if err != nil {
			return 1
		}
	case "login":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "token required")
			return 1
		}
		token := args[1]

		err := SaveAuth(authPath, Auth{Token: token})
		if err != nil {
			os.Remove(authPath)
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	case "clearcache":
		os.Remove(cachePath)
	case "logout":
		os.Remove(authPath)
	default:
		fmt.Fprintln(os.Stderr, "invalid action")
		return 1
	}

	return 0
}
