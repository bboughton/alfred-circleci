package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bboughton/alfred-circleci/alfred"
	"github.com/bboughton/alfred-circleci/cli"
)

type Auth struct {
	Token string
}

func LoadAuth(path string) Auth {
	var auth Auth
	file, err := os.Open(path)
	if err != nil {
		return auth
	}

	json.NewDecoder(file).Decode(&auth)
	return auth
}

func SaveAuth(path string, auth Auth) error {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if os.IsNotExist(err) {
		file, err = os.Create(path)
	}

	if err != nil {
		return err
	}

	if err = json.NewEncoder(file).Encode(auth); err != nil {
		return err
	}

	return nil
}

type AuthHandler struct {
	Auth    Auth
	Handler cli.NewHandler
}

func (h AuthHandler) Exec(out cli.OutputWriter, in *cli.Input) {
	args := in.Args
	resp := alfred.NewResponse()
	if h.Auth.Token == "" {
		var token string
		if len(args) > 1 {
			token = args[1]
		}
		resp.AddItem(alfred.NewItem("/login", "login "+token))
		err := alfred.WriteResponse(os.Stdout, resp)
		if err != nil {
			fmt.Println(err)
			out.ExitWith(1)
		}
	} else {
		h.Handler.Exec(out, in)
	}
}
