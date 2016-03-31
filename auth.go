package main

import (
	"encoding/json"
	"os"
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
